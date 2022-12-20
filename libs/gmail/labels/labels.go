package labels

import (
	"google.golang.org/api/gmail/v1"
)

// Labels contains all labels managed by @SRC.
type Labels struct {
	// SRC is the @SRC parent label for all @SRC labels.
	SRC *gmail.Label
	// Jobs is the @SRC/Jobs folder for all @SRC job labels.
	Jobs *gmail.Label
	// JobOpportunity is the @SRC/Jobs/Opportunity label for identified job opportunities.
	JobsOpportunity *gmail.Label
	// Allow is the @SRC/Allow folder for senders and domains that always bypass SRC
	Allow *gmail.Label
	// AllowSender is the @SRC/Allow/Sender folder for senders that always bypass SRC
	AllowSender *gmail.Label
	// AllowDomain is the @SRC/Allow/Domain folder for email domains that always bypass SRC
	AllowDomain *gmail.Label
	// Block is the @SRC/Allow folder for senders and domains that are automatically removed from your inbox.
	// Blocked emails are not analyzed by SRC.
	Block *gmail.Label
	// BlockSender is the @SRC/Block/Sender folder for senders that are automatically removed from your inbox.
	// Blocked emails are not analyzed by SRC.
	BlockSender *gmail.Label
	// BlockDomain is the @SRC/Block/Domain folder for email domains that are automatically removed from your inbox.
	// Blocked emails are not analyzed by SRC.
	BlockDomain *gmail.Label
	// BlockGraveyard is the @SRC/Block/Graveyard folder for emails that have been removed from your inbox.
	// Blocked emails are not be analyzed by SRC.
	BlockGraveyard *gmail.Label
}

var (
	SRC = &gmail.Label{
		Name:                "@SRC",
		LabelListVisibility: "labelShow",
		// Hide folder labels
		MessageListVisibility: "hide",
		Color: &gmail.LabelColor{
			BackgroundColor: "#4986e7",
			TextColor:       "#ffffff",
		},
	}
	Jobs = &gmail.Label{
		Name: SRC.Name + "/Jobs",
		// Hide folder labels
		MessageListVisibility: "hide",
		Color:                 SRC.Color,
	}
	JobsOpportunity = &gmail.Label{
		Name: Jobs.Name + "/Opportunity",
		// Show leaf labels
		MessageListVisibility: "show",
		// use same color as parent
		Color: Jobs.Color,
	}
	Allow = &gmail.Label{
		Name: SRC.Name + "/Allow",
		// Hide folder labels
		MessageListVisibility: "hide",
		Color: &gmail.LabelColor{
			BackgroundColor: "#16a765",
			TextColor:       "#ffffff",
		},
	}
	AllowSender = &gmail.Label{
		Name: Allow.Name + "/Sender",
		// Show leaf labels
		MessageListVisibility: "show",
		Color:                 Allow.Color,
	}
	AllowDomain = &gmail.Label{
		Name: Allow.Name + "/Domain",
		// Show leaf labels
		MessageListVisibility: "show",
		Color:                 Allow.Color,
	}
	Block = &gmail.Label{
		Name: SRC.Name + "/Block",
		// Hide folder labels
		MessageListVisibility: "hide",
		Color: &gmail.LabelColor{
			BackgroundColor: "#fb4c2f",
			TextColor:       "#ffffff",
		},
	}
	BlockSender = &gmail.Label{
		Name:                  Block.Name + "/Sender",
		MessageListVisibility: "show",
		Color:                 Block.Color,
	}
	BlockDomain = &gmail.Label{
		Name:                  Block.Name + "/Domain",
		MessageListVisibility: "show",
		Color:                 Block.Color,
	}
	BlockGraveyard = &gmail.Label{
		Name:                  Block.Name + "/Graveyard",
		MessageListVisibility: "show",
		Color:                 Block.Color,
	}
)
