package hasura

import (
	"reflect"
	"sort"
	"testing"
)

func TestCmd_setCommand(t *testing.T) {
	t.Parallel()

	type fields struct {
		called  string
		command []string
		files   []fileInfo
		options map[string]any
		target  string
	}

	tests := map[string]struct {
		fields fields
		want   *HasuraCmd
	}{
		"case: not set h.target": {
			fields: fields{
				called: "seed apply",
			},
			want: &HasuraCmd{
				called:  "seed apply",
				command: []string{"seed", "apply"},
			},
		},
		"case: h.called == 'seed apply'": {
			fields: fields{
				called: "seed apply",
				target: "xxx.sql",
			},
			want: &HasuraCmd{
				called:      "seed apply",
				applyTarget: "xxx.sql",
				command:     []string{"seed", "apply", "--file", "xxx.sql"},
			},
		},
		"case: h.called == 'seed apply', set string options": {
			fields: fields{
				called: "seed apply",
				target: "xxx.sql",
				options: map[string]any{
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:      "seed apply",
				applyTarget: "xxx.sql",
				command:     []string{"seed", "apply", "--file", "xxx.sql", "--admin-secret", "secret"},
				options: map[string]any{
					"admin-secret": "secret",
				},
			},
		},
		"case: h.called == 'seed apply', set bool options": {
			fields: fields{
				called: "seed apply",
				target: "xxx.sql",
				options: map[string]any{
					"no-color": false,
				},
			},
			want: &HasuraCmd{
				called:      "seed apply",
				applyTarget: "xxx.sql",
				command:     []string{"seed", "apply", "--file", "xxx.sql", "--no-color", "false"},
				options: map[string]any{
					"no-color": false,
				},
			},
		},
		"case: h.called == 'seed apply', set bool and string options": {
			fields: fields{
				called: "seed apply",
				target: "xxx.sql",
				options: map[string]any{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:      "seed apply",
				applyTarget: "xxx.sql",
				command:     []string{"seed", "apply", "--file", "xxx.sql", "--admin-secret", "secret", "--no-color", "false"},
				options: map[string]any{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
		},
		"case: h.called != 'seed apply'": {
			fields: fields{
				called: "migrate",
				target: "xxx.sql",
			},
			want: &HasuraCmd{
				called:      "migrate",
				applyTarget: "xxx.sql",
			},
		},
		"case: h.called == 'migrate apply'": {
			fields: fields{
				called: "migrate apply",
				target: "xxx",
			},
			want: &HasuraCmd{
				called:      "migrate apply",
				applyTarget: "xxx",
				command:     []string{"migrate", "apply", "--version", "xxx"},
			},
		},
		"case: h.called == 'migrate apply', set string options": {
			fields: fields{
				called: "migrate apply",
				target: "xxx",
				options: map[string]any{
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:      "migrate apply",
				applyTarget: "xxx",
				command:     []string{"migrate", "apply", "--version", "xxx", "--admin-secret", "secret"},
				options: map[string]any{
					"admin-secret": "secret",
				},
			},
		},
		"case: h.called == 'migrate apply', set bool options": {
			fields: fields{
				called: "migrate apply",
				target: "xxx",
				options: map[string]any{
					"no-color": false,
				},
			},
			want: &HasuraCmd{
				called:      "migrate apply",
				applyTarget: "xxx",
				command:     []string{"migrate", "apply", "--version", "xxx", "--no-color", "false"},
				options: map[string]any{
					"no-color": false,
				},
			},
		},
		"case: h.called == 'migrate', set bool and string options": {
			fields: fields{
				called: "migrate apply",
				target: "xxx",
				options: map[string]any{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:      "migrate apply",
				applyTarget: "xxx",
				command:     []string{"migrate", "apply", "--version", "xxx", "--admin-secret", "secret", "--no-color", "false"},
				options: map[string]any{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
		},
		"case: h.called == 'migrate delete'": {
			fields: fields{
				called: "migrate delete",
				target: "xxx",
			},
			want: &HasuraCmd{
				called:      "migrate delete",
				applyTarget: "xxx",
				command:     []string{"migrate", "delete", "--version", "xxx"},
			},
		},
		"case: h.called == 'migrate delete', set string options": {
			fields: fields{
				called: "migrate delete",
				target: "xxx",
				options: map[string]any{
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:      "migrate delete",
				applyTarget: "xxx",
				command:     []string{"migrate", "delete", "--version", "xxx", "--admin-secret", "secret"},
				options: map[string]any{
					"admin-secret": "secret",
				},
			},
		},
		"case: h.called == 'migrate delete', set bool options": {
			fields: fields{
				called: "migrate delete",
				target: "xxx",
				options: map[string]any{
					"no-color": false,
				},
			},
			want: &HasuraCmd{
				called:      "migrate delete",
				applyTarget: "xxx",
				command:     []string{"migrate", "delete", "--version", "xxx", "--no-color", "false"},
				options: map[string]any{
					"no-color": false,
				},
			},
		},
		"case: h.called == 'migrate delete', set bool and string options": {
			fields: fields{
				called: "migrate delete",
				target: "xxx",
				options: map[string]any{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:      "migrate delete",
				applyTarget: "xxx",
				command:     []string{"migrate", "delete", "--version", "xxx", "--admin-secret", "secret", "--no-color", "false"},
				options: map[string]any{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
		},
		"case: h.called != 'seed apply' & h.called != 'migrate apply & h.called != 'migrate delete'": {
			fields: fields{
				called: "hoge",
				target: "xxx",
			},
			want: &HasuraCmd{
				called:      "hoge",
				applyTarget: "xxx",
			},
		},
	}
	for testCase, tt := range tests {
		tt := tt
		t.Run(testCase, func(t *testing.T) {
			t.Parallel()
			h := &HasuraCmd{
				called:      tt.fields.called,
				command:     tt.fields.command,
				files:       tt.fields.files,
				options:     tt.fields.options,
				applyTarget: tt.fields.target,
			}
			got := h.setCommand()
			sort.Strings(got.command)
			sort.Strings(tt.want.command)

			if !reflect.DeepEqual(got.called, tt.want.called) {
				t.Errorf("Cmd.setCommand() = %v, want.called %v", got.called, tt.want.called)
			}
			if !reflect.DeepEqual(got.applyTarget, tt.want.applyTarget) {
				t.Errorf("HasuraCmd.setCommand() = %v, want.target %v", got.applyTarget, tt.want.applyTarget)
			}
			if !reflect.DeepEqual(got.command, tt.want.command) {
				t.Errorf("Cmd.setCommand() = %v, want.command %v", got.command, tt.want.command)
			}
			if !reflect.DeepEqual(got.files, tt.want.files) {
				t.Errorf("HasuraCmd.setCommand() = %v, want.files %v", got.files, tt.want.files)
			}
			if !reflect.DeepEqual(got.options, tt.want.options) {
				for _, set := range tt.want.options {
					var ok bool
					for _, gotSet := range got.options {
						if reflect.DeepEqual(set, gotSet) {
							ok = true

							break
						}
					}
					if !ok {
						t.Errorf("Cmd.setCommand() = %v, want.options %v", got.options, tt.want.options)
					}
				}
				if len(tt.want.options) != len(got.options) {
					t.Errorf("Cmd.setCommand() = %v, want.options %v", got.options, tt.want.options)
				}
			}
		})
	}
}
