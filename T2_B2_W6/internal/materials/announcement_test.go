package materials

import (
	"errors"
	"publish/internal/publicate"
	"slices"
	"testing"
)

func TestCreateAnnouncement(t *testing.T) {
	tests := []struct {
		name               string
		announcementName   string
		description        string
		header             string
		contentDescription string
		price              int
		ID                 int
		announcements      map[int]publicate.Material
		wantAnnouncements  map[int]publicate.Material
		wantErr            error
	}{
		{
			name:               "succesfull case",
			announcementName:   "vlad",
			description:        "finally learned go",
			header:             "LEARNING IS POWER",
			contentDescription: "Go school",
			price:              30,
			ID:                 20,
			announcements: map[int]publicate.Material{
				10: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},

			wantAnnouncements: map[int]publicate.Material{
				10: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},

				20: &Announcement{
					Profile: Profile{
						Name:        "vlad",
						Description: "finally learned go",
					},
					Material: Material{
						Header:    "LEARNING IS POWER",
						Status:    "draft",
						ActionLog: []string{"Announcemend created successfully!"},
					},
					Content: Content{
						Price:       30,
						Description: "Go school",
					},
				},
			},
			wantErr: nil,
		},

		{
			name:               "exist ID",
			announcementName:   "vlad",
			description:        "finally learned go",
			header:             "LEARNING IS POWER",
			contentDescription: "Go school",
			price:              30,
			ID:                 20,
			announcements: map[int]publicate.Material{
				20: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},

			wantAnnouncements: map[int]publicate.Material{
				20: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},
			wantErr: errIdExist,
		},

		{
			name:               "empty name",
			announcementName:   "",
			description:        "finally learned go",
			header:             "LEARNING IS POWER",
			contentDescription: "Go school",
			price:              30,
			ID:                 20,
			announcements: map[int]publicate.Material{
				10: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},

			wantAnnouncements: map[int]publicate.Material{
				10: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},
			wantErr: errNameEmpty,
		},

		{
			name:               "empty header",
			announcementName:   "vlad",
			description:        "finally learned go",
			header:             "",
			contentDescription: "Go school",
			price:              30,
			ID:                 20,
			announcements: map[int]publicate.Material{
				10: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},

			wantAnnouncements: map[int]publicate.Material{
				10: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},
			wantErr: errHeaderEmpty,
		},

		{
			name:               "invalid price",
			announcementName:   "vlad",
			description:        "finally learned go",
			header:             "LEARNING IS POWER",
			contentDescription: "Go school",
			price:              -20,
			ID:                 10,
			announcements: map[int]publicate.Material{
				20: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},

			wantAnnouncements: map[int]publicate.Material{
				20: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},
			wantErr: errPriceInvalid,
		},

		{
			name:               "invalid price",
			announcementName:   "vlad",
			description:        "finally learned go",
			header:             "LEARNING IS POWER",
			contentDescription: "Go school",
			price:              -20,
			ID:                 10,
			announcements: map[int]publicate.Material{
				20: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},

			wantAnnouncements: map[int]publicate.Material{
				20: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},
			wantErr: errPriceInvalid,
		}, {
			name:               "edge price case",
			announcementName:   "vlad",
			description:        "finally learned go",
			header:             "LEARNING IS POWER",
			contentDescription: "Go school",
			price:              0,
			ID:                 10,
			announcements: map[int]publicate.Material{
				20: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},
			},

			wantAnnouncements: map[int]publicate.Material{
				20: &Announcement{
					Profile: Profile{
						Name:        "dime",
						Description: "dima krasavec",
					},
					Material: Material{
						Header:    "somethink",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Content: Content{
						Price:       90,
						Description: "Sell a car",
					},
				},

				10: &Announcement{
					Profile: Profile{
						Name:        "vlad",
						Description: "finally learned go",
					},
					Material: Material{
						Header:    "LEARNING IS POWER",
						Status:    "draft",
						ActionLog: []string{"Announcemend created successfully!"},
					},
					Content: Content{
						Price:       0,
						Description: "Go school",
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateAnnouncement(tt.announcementName, tt.description, tt.header, tt.contentDescription, tt.price, tt.ID, tt.announcements)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}
			if len(tt.announcements) == 0 {
				t.Fatal("announcement is empty")
			}

			if len(tt.announcements) != len(tt.wantAnnouncements) {
				t.Fatalf("expected %v lenth of maps, got %v", len(tt.wantAnnouncements), len(tt.announcements))
			}

			for id := range tt.announcements {
				a := *tt.announcements[id].(*Announcement)
				wantA := *tt.wantAnnouncements[id].(*Announcement)

				if !slices.Equal(a.ActionLog, wantA.ActionLog) {
					t.Errorf("expected action log %v, got %v", wantA.ActionLog, a.ActionLog)
				}

				if a.Name != wantA.Name || a.Profile.Description != wantA.Profile.Description {
					t.Errorf("expected announcement profile %v, got %v", wantA.Profile, a.Profile)
				}

				if a.Header != wantA.Header || a.Status != wantA.Status {
					t.Errorf("expected material %v, got %v", wantA.Material, a.Material)
				}

				if a.Price != wantA.Price || a.Content.Description != wantA.Content.Description {
					t.Errorf("expected content %v, got %v", wantA.Content, a.Content)
				}
			}
		})
	}
}

func TestAnnouncementEdit(t *testing.T) {
	tests := []struct {
		name             string
		announName       string
		header           string
		description      string
		price            int
		announcement     Announcement
		wantAnnouncement Announcement
		wantLog          []string
		wantErr          error
	}{
		{
			name:        "success case",
			announName:  "",
			header:      "success",
			description: "hello",
			price:       10,
			announcement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantAnnouncement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header: "success",
					Status: "draft",
				},
				Content: Content{
					Description: "hello",
					Price:       10,
				},
			},

			wantLog: []string{"success", "header description Edited successfully!"},
			wantErr: nil,
		},

		{
			name:        "published status case",
			announName:  "dima",
			header:      "success",
			description: "hello",
			price:       10,

			announcement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantAnnouncement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantErr: ErrStatusPublished,
			wantLog: []string{"success"},
		},

		{
			name:        "empty data case",
			announName:  "",
			header:      "",
			description: "",
			price:       -1,

			announcement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantAnnouncement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantErr: errEditDataEmpty,
			wantLog: []string{"success"},
		},

		{
			name:        "invalid price",
			announName:  "Blad",
			header:      "Successful decision",
			description: "haha",
			price:       -4,

			announcement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantAnnouncement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantErr: errPriceInvalid,
			wantLog: []string{"success"},
		},

		{
			name:        "edge price case",
			announName:  "dima",
			header:      "",
			description: "",
			price:       -1,
			announcement: Announcement{
				Profile: Profile{
					Name:        "Vlad",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantAnnouncement: Announcement{
				Profile: Profile{
					Name:        "dima",
					Description: "16 yo",
				},
				Material: Material{
					Header:    "Learning go",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Description: "Bought a car",
					Price:       20,
				},
			},

			wantLog: []string{"success", "name Edited successfully!"},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.announcement.Edit(tt.announName, tt.header, tt.description, tt.price)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if tt.announcement.Profile != tt.wantAnnouncement.Profile {
				t.Errorf("expected profile %v, got %v", tt.announcement.Profile, tt.wantAnnouncement.Profile)
			}

			if tt.announcement.Header != tt.wantAnnouncement.Header {
				t.Errorf("expected header %v, got %v", tt.wantAnnouncement.Header, tt.announcement.Header)
			}

			if tt.announcement.Status != tt.wantAnnouncement.Status {
				t.Errorf("expected status %v, got %v", tt.wantAnnouncement, tt.announcement)
			}

			if tt.announcement.Content != tt.wantAnnouncement.Content {
				t.Errorf("expected content %v, got %v", tt.wantAnnouncement.Content, tt.announcement.Content)
			}

			if !slices.Equal(tt.announcement.ActionLog, tt.wantLog) {
				t.Errorf("expected action log %v, got %v", tt.announcement.ActionLog, tt.wantLog)
			}
		})
	}
}

func TestAnnouncementPublicate(t *testing.T) {
	tests := []struct {
		name             string
		announcement     Announcement
		wantAnnouncement Announcement
		wantErr          error
	}{
		{
			name: "success case",
			announcement: Announcement{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Price:       20,
					Description: "Successfull",
				},
			},

			wantAnnouncement: Announcement{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Price:       20,
					Description: "Successfull",
				},
			},

			wantErr: nil,
		},

		{
			name: "not enough description",
			announcement: Announcement{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Price:       20,
					Description: "",
				},
			},

			wantAnnouncement: Announcement{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Price:       20,
					Description: "",
				},
			},

			wantErr: ErrContentEmpty,
		},

		{
			name: "published status",
			announcement: Announcement{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Price:       20,
					Description: "Successfull",
				},
			},

			wantAnnouncement: Announcement{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Content: Content{
					Price:       20,
					Description: "Successfull",
				},
			},

			wantErr: ErrStatusPublished,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.announcement.Publicate()

			if err != tt.wantErr {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if tt.announcement.Profile != tt.wantAnnouncement.Profile {
				t.Errorf("expected profile %v, got %v", tt.wantAnnouncement.Profile, tt.announcement.Profile)
			}

			if !slices.Equal(tt.announcement.ActionLog, tt.wantAnnouncement.ActionLog) {
				t.Errorf("expected action log %v, got %v", tt.announcement.ActionLog, tt.wantAnnouncement.ActionLog)
			}

			if tt.announcement.Content != tt.wantAnnouncement.Content {
				t.Errorf("expected content %v, got %v", tt.wantAnnouncement.Content, tt.announcement.Content)
			}

			if tt.announcement.Header != tt.wantAnnouncement.Header || tt.announcement.Status != tt.wantAnnouncement.Status {
				t.Errorf("expected material %v, got %v", tt.wantAnnouncement.Material, tt.announcement.Material)
			}
		})
	}
}
