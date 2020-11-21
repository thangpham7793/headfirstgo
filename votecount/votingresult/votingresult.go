//Package votingresult mimics oop and redux (with init and instance methods)
package votingresult

import (
	"bufio"
	"errors"
	"fmt"
	"general"
	"os"
	"sort"
)

//VotingResult defines props. Struct is like model (MongoDB, etc.) but shouldn't be tied to a particular storage implementation?
type VotingResult struct {
	votes map[string]int
}

//New inits
func New() *VotingResult {
	return &VotingResult{make(map[string]int)}
}

//PrintCandidateNames prints all name
func (r *VotingResult) PrintCandidateNames() {
	for n := range r.votes {
		fmt.Println(n)
	}
}

//DisqualifyCandidate removes candidate from the votes map
func (r *VotingResult) DisqualifyCandidate(name string) {
	if len(name) == 0 {
		general.HandleErr(errors.New("no name given"))
	}
	delete(r.votes, name)
}

//EnrollCandidate add candidates to the votes map
func (r *VotingResult) EnrollCandidate(name string) {
	if len(name) == 0 {
		general.HandleErr(errors.New("no name given"))
	}
	r.votes[name] = 0
}

//PrintOneCandidateVoteCount prints the name and vote count of one candidate
func (r *VotingResult) PrintOneCandidateVoteCount(name string) {

	if len(name) == 0 {
		general.HandleErr(errors.New("no name given"))
	}

	for n, v := range r.votes {
		if n == name {
			fmt.Printf("Candidate %s: %d votes\n", n, v)
			return
		}
	}
	general.HandleErr(errors.New("no such name found"))
}

//PrintNamesAndResults prints the name and vote count of one candidate
func (r *VotingResult) PrintNamesAndResults() {

	names := make([]string, 0)
	for n := range r.votes {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, n := range names {
		fmt.Printf("Candidate %s: %d votes\n", n, r.votes[n])
	}
}

//IsCandidateEnrolled checks if a candidate is enrolled
func (r *VotingResult) IsCandidateEnrolled(name string) bool {
	if len(name) == 0 {
		general.HandleErr(errors.New("no name given"))
	}
	for n := range r.votes {
		if n == name {
			return true
		}
	}
	return false
}

//CountVotes prints vote counts for each candidate
func (r *VotingResult) CountVotes(f *os.File) {

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		name := sc.Text()
		r.votes[name]++
	}

	general.HandleErr(sc.Err())

	for k, v := range r.votes {
		fmt.Printf("Candidate %s: %d votes\n", k, v)
	}
}
