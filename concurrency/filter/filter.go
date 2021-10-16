package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

var workers = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(workers)
	log.SetFlags(0)
	flags := handleCommandLine()
	filterFiles(flags)
}

type Job struct {
	filename     string
	modifiedTime time.Time
	resultsCh    chan<- Results
}

func (job *Job) FilterByDate(flags *LineFlags) <-chan Job {
	out := make(chan Job)
	go func() {
		modifiedTime := job.modifiedTime
		yy, mm, dd := modifiedTime.Date()
		switch {
		case flags.day == 0 && flags.year == 0 && flags.month == 0:
			out <- Job{job.filename, job.modifiedTime, job.resultsCh}
		case flags.day == 0 && flags.month == 0 && flags.year == yy:
			out <- Job{job.filename, job.modifiedTime, job.resultsCh}
		case flags.day == 0 && flags.year == 0 && flags.month == int(mm):
			out <- Job{job.filename, job.modifiedTime, job.resultsCh}
		case flags.month == 0 && flags.year == 0 && flags.day == dd:
			out <- Job{job.filename, job.modifiedTime, job.resultsCh}
		case flags.day == 0 && flags.month == int(mm) && flags.year == yy:
			out <- Job{job.filename, job.modifiedTime, job.resultsCh}
		case flags.month == 0 && flags.day == dd && flags.year == yy:
			out <- Job{job.filename, job.modifiedTime, job.resultsCh}
		case flags.year == 0 && flags.day == dd && flags.month == int(mm):
			out <- Job{job.filename, job.modifiedTime, job.resultsCh}
		case flags.year == yy && flags.day == dd && flags.month == int(mm):
			out <- Job{job.filename, job.modifiedTime, job.resultsCh}
		}
		close(out)
	}()
	return out
}
func (job *Job) FilterBySuffix(suffixes []string, in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for item := range in {
			if len(suffixes) == 0 { // any file accepted
				out <- Job{item.filename, item.modifiedTime, item.resultsCh}
				continue
			}
			fileExt := strings.ToLower(filepath.Ext(item.filename))
			for _, suffix := range suffixes {
				if suffix == fileExt {
					out <- Job{item.filename, item.modifiedTime, item.resultsCh}
				}
			}
		}
		close(out)
	}()
	return out
}

type Results struct {
	filename     string
	modifiedTime time.Time
}

type LineFlags struct {
	day      int
	month    int
	year     int
	files    []string
	suffixes []string
}

func handleCommandLine() (flags *LineFlags) {
	flags = &LineFlags{}
	flag.IntVar(&flags.day, "d", 0, "day number ( <= 0 means any day)")
	flag.IntVar(&flags.month, "m", 0, "month number ( <= 0 means any month)")
	flag.IntVar(&flags.year, "y", 0, "year number ( <= 0 means any year)")
	suffixes := flag.String("s", "", "file extension suffix separated by ',' (any file by default)")
	flag.Parse()

	flags.files = flag.Args()
	if flags.day > 31 || flags.month > 12 {
		log.Fatalln("error: parse date: 'day' <= 31, 'month' <= 12")
	}

	if *suffixes != "" {
		flags.suffixes = strings.Split(strings.ToLower(*suffixes), ",")
	}
	return flags
}

func filterFiles(flags *LineFlags) {
	jobCh := make(chan Job, workers)
	resultCh := make(chan Results, len(flags.files))
	doneCh := make(chan struct{}, workers)
	go addJobs(jobCh, flags.files, resultCh)
	for i := 0; i < workers; i++ {
		go doJob(doneCh, flags, jobCh)
	}
	go waitUntilComplete(doneCh, resultCh)
	printResults(resultCh)
}

func addJobs(jobCh chan<- Job, files []string, resultCh chan<- Results) {
	for _, filename := range files {
		fileInfo, err := os.Stat(filename)
		if err != nil {
			continue
		}
		modifiedTime := fileInfo.ModTime()
		jobCh <- Job{filename, modifiedTime, resultCh}
	}
	close(jobCh)
}

func doJob(doneCh chan<- struct{}, flags *LineFlags, jobCh <-chan Job) {
	for job := range jobCh {
		dateFilterCh := job.FilterByDate(flags)
		suffixFilterCh := job.FilterBySuffix(flags.suffixes, dateFilterCh)
		for jobs := range suffixFilterCh {
			jobs.resultsCh <- Results{jobs.filename, jobs.modifiedTime}
		}
	}
	doneCh <- struct{}{}
}

func waitUntilComplete(doneCh <-chan struct{}, resultCh chan Results) {
	for i := 0; i < workers; i++ {
		<-doneCh
	}
	close(resultCh)
}

func printResults(results <-chan Results) {
	rSlice := make([]Results, 0)

	for result := range results {
		rSlice = append(rSlice, result)
	}

	sort.Slice(rSlice, func(i, j int) bool {
		return rSlice[i].modifiedTime.After(rSlice[j].modifiedTime)
	})
	for _, res := range rSlice {
		yy, mm, dd := res.modifiedTime.Date()
		h, m := res.modifiedTime.Hour(), res.modifiedTime.Minute()
		_, fileName := filepath.Split(res.filename)
		fmt.Printf("modified: %.2d-%.2d-%.2d %.2d:%.2d %s\n", dd, int(mm), yy, h, m,
			fileName)
	}
}
