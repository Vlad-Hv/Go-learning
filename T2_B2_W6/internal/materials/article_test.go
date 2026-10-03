package materials

import (
	"errors"
	"publish/internal/publicate"
	"slices"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	tests := []struct {
		name         string
		articleName  string
		description  string
		header       string
		context      string
		ID           int
		articles     map[int]publicate.Material
		wantArticles map[int]publicate.Material
		wantErr      error
	}{
		{
			name:        "successfull case",
			articleName: "vlad",
			description: "learned go",
			header:      "DONE IT",
			context:     "finally learned go",
			ID:          20,
			articles: map[int]publicate.Material{
				10: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "yes",
					},
					Material: Material{
						Header:    "YES",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Context: "finally learned something",
				},
			},
			wantArticles: map[int]publicate.Material{
				10: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "yes",
					},
					Material: Material{
						Header:    "YES",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Context: "finally learned something",
				},

				20: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "learned go",
					},
					Material: Material{
						Header:    "DONE IT",
						Status:    "draft",
						ActionLog: []string{"Article created succesfully"},
					},
					Context: "finally learned go",
				},
			},
			wantErr: nil,
		},

		{
			name:        "already exist ID",
			articleName: "vlad",
			description: "learned go",
			header:      "DONE IT",
			context:     "finally learned go",
			ID:          10,
			articles: map[int]publicate.Material{
				10: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "yes",
					},
					Material: Material{
						Header:    "YES",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Context: "finally learned something",
				},
			},
			wantArticles: map[int]publicate.Material{
				10: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "yes",
					},
					Material: Material{
						Header:    "YES",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Context: "finally learned something",
				},
			},
			wantErr: errIdExist,
		},

		{
			name:        "empty name case",
			articleName: "",
			description: "learned go",
			header:      "DONE IT",
			context:     "finally learned go",
			ID:          20,
			articles: map[int]publicate.Material{
				10: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "yes",
					},
					Material: Material{
						Header:    "YES",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Context: "finally learned something",
				},
			},
			wantArticles: map[int]publicate.Material{
				10: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "yes",
					},
					Material: Material{
						Header:    "YES",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Context: "finally learned something",
				},
			},
			wantErr: errNameEmpty,
		},

		{
			name:        "empty header case",
			articleName: "vlad",
			description: "learned go",
			header:      "",
			context:     "finally learned go",
			ID:          40,
			articles: map[int]publicate.Material{
				10: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "yes",
					},
					Material: Material{
						Header:    "YES",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Context: "finally learned something",
				},
			},
			wantArticles: map[int]publicate.Material{
				10: &Article{
					Profile: Profile{
						Name:        "vlad",
						Description: "yes",
					},
					Material: Material{
						Header:    "YES",
						Status:    "draft",
						ActionLog: []string{"success"},
					},
					Context: "finally learned something",
				},
			},
			wantErr: errHeaderEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateArticle(tt.articleName, tt.description, tt.header, tt.context, tt.ID, tt.articles)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if len(tt.articles) != len(tt.wantArticles) {
				t.Fatalf("expected %v lenth of maps, got %v", len(tt.wantArticles), len(tt.articles))
			}

			for id := range tt.wantArticles {
				article, ok := tt.articles[id].(*Article)
				if !ok {
					t.Fatal("expected ok")
				}
				wantArticle := *tt.wantArticles[id].(*Article)

				if !slices.Equal(article.ActionLog, wantArticle.ActionLog) {
					t.Errorf("expected slice %v, got %v", wantArticle.ActionLog, article.ActionLog)
				}

				if article.Name != wantArticle.Name || article.Description != wantArticle.Description {
					t.Errorf("expected article profile %v, got %v", wantArticle.Profile, article.Profile)
				}

				if article.Header != wantArticle.Header || article.Status != wantArticle.Status {
					t.Errorf("expected material %v, got %v", wantArticle.Material, article.Material)
				}

				if article.Context != wantArticle.Context {
					t.Errorf("expected context %v, got %v", wantArticle.Context, article.Context)
				}
			}
		})
	}
}

//successfull case, empty name, empty description, already exist ID, empty header

func TestAcrticleEdit(t *testing.T) {
	tests := []struct {
		name        string
		articleName string
		header      string
		context     string
		article     Article
		wantArticle Article
		wantErr     error
	}{
		{
			name:        "success case",
			articleName: "dima",
			header:      "learned",
			context:     "super",
			article: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "found",
				},
				Material: Material{
					Header:    "found",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Context: "string",
			},
			wantArticle: Article{
				Profile: Profile{
					Name:        "dima",
					Description: "found",
				},
				Material: Material{
					Header:    "learned",
					Status:    "draft",
					ActionLog: []string{"success", "name header context Edited succesfully"},
				},
				Context: "super",
			},

			wantErr: nil,
		},

		{
			name:        "published status case",
			articleName: "dima",
			header:      "learned",
			context:     "super",
			article: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "found",
				},
				Material: Material{
					Header:    "found",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "string",
			},
			wantArticle: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "found",
				},
				Material: Material{
					Header:    "found",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "string",
			},

			wantErr: ErrStatusPublished,
		},

		{
			name:        "published status case",
			articleName: "",
			header:      "",
			context:     "",
			article: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "found",
				},
				Material: Material{
					Header:    "Hello",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Context: "string",
			},
			wantArticle: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "found",
				},
				Material: Material{
					Header:    "Hello",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Context: "string",
			},

			wantErr: errEditDataEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.article.Edit(tt.articleName, tt.header, tt.context, 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if tt.article.Profile != tt.wantArticle.Profile {
				t.Errorf("expected profile %v, got %v", tt.article.Profile, tt.wantArticle.Profile)
			}

			if tt.article.Header != tt.wantArticle.Header {
				t.Errorf("expected header %v, got %v", tt.wantArticle.Header, tt.article.Header)
			}

			if tt.article.Status != tt.wantArticle.Status {
				t.Errorf("expected status %v, got %v", tt.wantArticle, tt.article)
			}

			if tt.article.Context != tt.wantArticle.Context {
				t.Errorf("expected content %v, got %v", tt.wantArticle.Context, tt.article.Context)
			}

			if !slices.Equal(tt.article.ActionLog, tt.wantArticle.ActionLog) {
				t.Errorf("expected action log %v, got %v", tt.article.ActionLog, tt.wantArticle.ActionLog)
			}
		})
	}
}

func TestArticlePublicate(t *testing.T) {
	tests := []struct {
		name        string
		article     Article
		wantArticle Article
		wantErr     error
	}{
		{
			name: "success case",
			article: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Context: "hello guis, I want to",
			},

			wantArticle: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "hello guis, I want to",
			},

			wantErr: nil,
		},

		{
			name: "not enough description",
			article: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Context: "hello i'm vlad 16 y",
			},

			wantArticle: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "draft",
					ActionLog: []string{"success"},
				},
				Context: "hello i'm vlad 16 y",
			},

			wantErr: ErrContextSmall,
		},

		{
			name: "published status",
			article: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "hello guis, I want to buy a new car, so let me do this",
			},

			wantArticle: Article{
				Profile: Profile{
					Name:        "vlad",
					Description: "pum pum pum",
				},
				Material: Material{
					Header:    "Learned",
					Status:    "published",
					ActionLog: []string{"success"},
				},
				Context: "hello guis, I want to buy a new car, so let me do this",
			},

			wantErr: ErrStatusPublished,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.article.Publicate()

			if err != tt.wantErr {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if tt.article.Profile != tt.wantArticle.Profile {
				t.Errorf("expected profile %v, got %v", tt.wantArticle.Profile, tt.article.Profile)
			}

			if !slices.Equal(tt.article.ActionLog, tt.wantArticle.ActionLog) {
				t.Errorf("expected action log %v, got %v", tt.article.ActionLog, tt.wantArticle.ActionLog)
			}

			if tt.article.Context != tt.wantArticle.Context {
				t.Errorf("expected profile %v, got %v", tt.wantArticle.Context, tt.article.Context)
			}

			if tt.article.Header != tt.wantArticle.Header || tt.article.Status != tt.wantArticle.Status {
				t.Errorf("expected material %v, got %v", tt.wantArticle.Material, tt.article.Material)
			}
		})
	}
}
