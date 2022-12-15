package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"
	"math/rand"

	"github.com/benbjohnson/immutable"
)

func (a *Team6Agent) LootActionNoProposal(baseAgent agent.BaseAgent) immutable.SortedMap[commons.ItemID, struct{}] {
	return *immutable.NewSortedMap[commons.ItemID, struct{}](nil)
}

func (a *Team6Agent) LootAction(
	baseAgent agent.BaseAgent,
	proposedLoot immutable.SortedMap[commons.ItemID, struct{}],
	acceptedProposal message.Proposal[decision.LootAction],
) immutable.SortedMap[commons.ItemID, struct{}] {
	return proposedLoot
}

func (a *Team6Agent) HandleLootInformation(m message.TaggedInformMessage[message.LootInform], agent agent.BaseAgent) {
}

func (a *Team6Agent) HandleLootRequest(m message.TaggedRequestMessage[message.LootRequest]) message.LootInform {
	//TODO implement me
	panic("implement me")
}

func (a *Team6Agent) HandleLootProposal(_ message.Proposal[decision.LootAction], _ agent.BaseAgent) decision.Intent {
	switch rand.Intn(3) {
	case 0:
		return decision.Positive
	case 1:
		return decision.Negative
	default:
		return decision.Abstain
	}
}

func (a *Team6Agent) HandleLootProposalRequest(proposal message.Proposal[decision.LootAction], _ agent.BaseAgent) bool {
	// TODO: Replace one of these with agent's own proposal
	return proposalSimilarity(proposal.Rules(), proposal.Rules()) > 0.6
}

func (a *Team6Agent) LootAllocation(baseAgent agent.BaseAgent, proposal message.Proposal[decision.LootAction], proposedAllocations immutable.Map[commons.ID, immutable.SortedMap[commons.ItemID, struct{}]]) immutable.Map[commons.ID, immutable.SortedMap[commons.ItemID, struct{}]] {
	return proposedAllocations
}
