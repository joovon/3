package main

type User struct {
	Jobs      []*Job
	FairShare float64 // Used-up compute time in the past (decays)
	nextPtr   int     // pointer suggesting next job to start. Reset on re-scan. len(Jobs) means no queued job
}

func NewUser() *User {
	return &User{}
}

//// used share in GPU*hours
//func (u *Share) Used() float64 {
//	used := u.used
//	for _, j := range u.Running {
//		used += j.Runtime().Seconds()
//	}
//	return used / 3600
//}

// nextJob looks for the next free job in the list.
// it does a tiny bit of linear search, starting from nextPtr.
//func (u *User) giveJob(node string) *Job {
//
//	return nil
//}

func (u *User) HasJob() bool {
	i := u.nextJobPtr()
	return i < len(u.Jobs)
}

func (u *User) nextJobPtr() int {
	for ; u.nextPtr < len(u.Jobs); u.nextPtr++ {
		j := u.Jobs[u.nextPtr]
		if j.IsQueued() {
			return u.nextPtr
		}
	}
	return u.nextPtr
}

//func (u *User) GiveJob(nodeAddr string) string {
//	var job *Job
//
//	// take random job from Queue map
//	for _, j := range u.Queue {
//		job = j
//		break
//	}
//	delete(u.Queue, job.URL)
//
//	job.Status = RUNNING
//	job.Start = time.Now()
//	job.Node = nodeAddr
//	job.Node = nodeAddr
//
//	u.Running[job.URL] = job
//
//	//log.Println("give job", job.URL, "->", nodeAddr)
//	return job.URL
//}
