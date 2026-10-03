package test

import (
	"errors"
	"publish/internal/materials"
	"publish/internal/publicate"
	"slices"
	"testing"
)

func TestPublicate(t *testing.T) {
	tests := []struct {
		name         string
		material     publicate.Material
		wantMaterial publicate.Material
		wantErr      error
	}{
		{
			name: "success announcement case",
			material: &materials.Announcement{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: materials.Content{
					Price:       20,
					Description: "Successfull",
				},
			},

			wantMaterial: &materials.Announcement{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success", "publicate done successfully"},
				},
				Content: materials.Content{
					Price:       20,
					Description: "Successfull",
				},
			},

			wantErr: nil,
		},

		{
			name: "success announcement case",
			material: &materials.Article{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Context: "hello so im too big and",
			},

			wantMaterial: &materials.Article{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success", "publicate done successfully"},
				},
				Context: "hello so im too big and",
			},

			wantErr: nil,
		},

		{
			name: "article status publish case",
			material: &materials.Article{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "hello so im too big and",
			},

			wantMaterial: &materials.Article{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "hello so im too big and",
			},

			wantErr: materials.ErrStatusPublished,
		},

		{
			name: "article small context case",
			material: &materials.Article{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "hello so im too big",
			},

			wantMaterial: &materials.Article{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "hello so im too big",
			},

			wantErr: materials.ErrContextSmall,
		},

		{
			name: "empty description case",
			material: &materials.Announcement{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: materials.Content{
					Price:       20,
					Description: "",
				},
			},

			wantMaterial: &materials.Announcement{
				Profile: materials.Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: materials.Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Content: materials.Content{
					Price:       20,
					Description: "",
				},
			},

			wantErr: materials.ErrContentEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := publicate.Publicate(tt.material)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			material, ok := tt.material.(*materials.Announcement)
			if !ok {
				material := tt.material.(*materials.Article)
				wantMaterial := tt.wantMaterial.(*materials.Article)
				if material.Profile != wantMaterial.Profile {
					t.Errorf("expected profile %v, got %v", wantMaterial.Profile, material.Profile)
				}

				if !slices.Equal(material.ActionLog, wantMaterial.ActionLog) {
					t.Errorf("expected action log %v, got %v", material.ActionLog, wantMaterial.ActionLog)
				}

				if material.Context != wantMaterial.Context {
					t.Errorf("expected content %v, got %v", wantMaterial.Context, material.Context)
				}

				if material.Header != wantMaterial.Header || material.Status != wantMaterial.Status {
					t.Errorf("expected material %v, got %v", wantMaterial.Material, material.Material)
				}

				return
			}

			wantMaterial, ok := tt.wantMaterial.(*materials.Announcement)
			if wantMaterial == nil {
				t.Fatal("expected want material exist")
			}

			if !ok {
				t.Fatal("expected ok")
			}

			if material.Profile != wantMaterial.Profile {
				t.Errorf("expected profile %v, got %v", wantMaterial.Profile, material.Profile)
			}

			if !slices.Equal(material.ActionLog, wantMaterial.ActionLog) {
				t.Errorf("expected action log %v, got %v", material.ActionLog, wantMaterial.ActionLog)
			}

			if material.Content != wantMaterial.Content {
				t.Errorf("expected content %v, got %v", wantMaterial.Content, material.Content)
			}

			if material.Header != wantMaterial.Header || material.Status != wantMaterial.Status {
				t.Errorf("expected material %v, got %v", wantMaterial.Material, material.Material)
			}
		})
	}
}
