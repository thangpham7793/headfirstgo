package main

import (
	"github.com/headfirstgo/general"
	"github.com/headfirstgo/votecount/utils"
	"github.com/headfirstgo/votecount/votingresult"
)

func main() {

	f := utils.ReadTextFile()

	defer func() {
		err := f.Close()
		general.HandleErr(err)
	}()

	r := votingresult.New()
	r.CountVotes(f)
	// r.PrintCandidateNames()
	// r.PrintOneCandidateVoteCount("A")
	// // r.PrintOneCandidateVoteCount("D")
	// r.DisqualifyCandidate("B")
	// r.PrintOneCandidateVoteCount("B")
	// r.EnrollCandidate("F")
	r.PrintNamesAndResults()
	// r.PrintOneCandidateVoteCount("F")
	// println(r.IsCandidateEnrolled("Z"))
}
