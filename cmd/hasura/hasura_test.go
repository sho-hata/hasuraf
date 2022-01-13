package hasura

import (
	"reflect"
	"sort"
	"testing"
)

func TestHasuraCmd_setCommand(t *testing.T) {
	t.Parallel()
	type fields struct {
		called    string
		command   []string
		fileNames []string
		options   map[string]interface{}
		target    string
	}
	tests := []struct {
		name   string
		fields fields
		want   *HasuraCmd
	}{
		{
			name: "case: not set h.target",
			fields: fields{
				called: "seed",
			},
			want: &HasuraCmd{
				called: "seed",
			},
		},
		{
			name: "case: h.called == 'seed'",
			fields: fields{
				called: "seed",
				target: "xxx.sql",
			},
			want: &HasuraCmd{
				called:  "seed",
				target:  "xxx.sql",
				command: []string{"seed", "apply", "--file", "xxx.sql"},
			},
		},
		{
			name: "case: h.called == 'seed', set string options",
			fields: fields{
				called: "seed",
				target: "xxx.sql",
				options: map[string]interface{}{
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:  "seed",
				target:  "xxx.sql",
				command: []string{"seed", "apply", "--file", "xxx.sql", "--admin-secret", "secret"},
				options: map[string]interface{}{
					"admin-secret": "secret",
				},
			},
		},
		{
			name: "case: h.called == 'seed', set bool options",
			fields: fields{
				called: "seed",
				target: "xxx.sql",
				options: map[string]interface{}{
					"no-color": false,
				},
			},
			want: &HasuraCmd{
				called:  "seed",
				target:  "xxx.sql",
				command: []string{"seed", "apply", "--file", "xxx.sql", "--no-color", "false"},
				options: map[string]interface{}{
					"no-color": false,
				},
			},
		},
		{
			name: "case: h.called == 'seed', set bool and string options",
			fields: fields{
				called: "seed",
				target: "xxx.sql",
				options: map[string]interface{}{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:  "seed",
				target:  "xxx.sql",
				command: []string{"seed", "apply", "--file", "xxx.sql", "--admin-secret", "secret", "--no-color", "false"},
				options: map[string]interface{}{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
		},
		{
			name: "case: h.called != 'seed'",
			fields: fields{
				called: "migrate",
				target: "xxx.sql",
			},
			want: &HasuraCmd{
				called: "migrate",
				target: "xxx.sql",
			},
		},
		{
			name: "case: h.called == 'migrate apply'",
			fields: fields{
				called: "migrate apply",
				target: "xxx",
			},
			want: &HasuraCmd{
				called:  "migrate apply",
				target:  "xxx",
				command: []string{"migrate", "apply", "--version", "xxx"},
			},
		},
		{
			name: "case: h.called == 'migrate apply', set string options",
			fields: fields{
				called: "migrate apply",
				target: "xxx",
				options: map[string]interface{}{
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:  "migrate apply",
				target:  "xxx",
				command: []string{"migrate", "apply", "--version", "xxx", "--admin-secret", "secret"},
				options: map[string]interface{}{
					"admin-secret": "secret",
				},
			},
		},
		{
			name: "case: h.called == 'migrate apply', set bool options",
			fields: fields{
				called: "migrate apply",
				target: "xxx",
				options: map[string]interface{}{
					"no-color": false,
				},
			},
			want: &HasuraCmd{
				called:  "migrate apply",
				target:  "xxx",
				command: []string{"migrate", "apply", "--version", "xxx", "--no-color", "false"},
				options: map[string]interface{}{
					"no-color": false,
				},
			},
		},
		{
			name: "case: h.called == 'migrate', set bool and string options",
			fields: fields{
				called: "migrate apply",
				target: "xxx",
				options: map[string]interface{}{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:  "migrate apply",
				target:  "xxx",
				command: []string{"migrate", "apply", "--version", "xxx", "--admin-secret", "secret", "--no-color", "false"},
				options: map[string]interface{}{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
		},
		{
			name: "case: h.called == 'migrate delete'",
			fields: fields{
				called: "migrate delete",
				target: "xxx",
			},
			want: &HasuraCmd{
				called:  "migrate delete",
				target:  "xxx",
				command: []string{"migrate", "delete", "--version", "xxx"},
			},
		},
		{
			name: "case: h.called == 'migrate delete', set string options",
			fields: fields{
				called: "migrate delete",
				target: "xxx",
				options: map[string]interface{}{
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:  "migrate delete",
				target:  "xxx",
				command: []string{"migrate", "delete", "--version", "xxx", "--admin-secret", "secret"},
				options: map[string]interface{}{
					"admin-secret": "secret",
				},
			},
		},
		{
			name: "case: h.called == 'migrate delete', set bool options",
			fields: fields{
				called: "migrate delete",
				target: "xxx",
				options: map[string]interface{}{
					"no-color": false,
				},
			},
			want: &HasuraCmd{
				called:  "migrate delete",
				target:  "xxx",
				command: []string{"migrate", "delete", "--version", "xxx", "--no-color", "false"},
				options: map[string]interface{}{
					"no-color": false,
				},
			},
		},
		{
			name: "case: h.called == 'migrate', set bool and string options",
			fields: fields{
				called: "migrate delete",
				target: "xxx",
				options: map[string]interface{}{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
			want: &HasuraCmd{
				called:  "migrate delete",
				target:  "xxx",
				command: []string{"migrate", "delete", "--version", "xxx", "--admin-secret", "secret", "--no-color", "false"},
				options: map[string]interface{}{
					"no-color":     false,
					"admin-secret": "secret",
				},
			},
		},
		{
			name: "case: h.called != 'seed' & h.called != 'migrate apply & h.called != 'migrate delete'",
			fields: fields{
				called: "hoge",
				target: "xxx",
			},
			want: &HasuraCmd{
				called: "hoge",
				target: "xxx",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := &HasuraCmd{
				called:    tt.fields.called,
				command:   tt.fields.command,
				fileNames: tt.fields.fileNames,
				options:   tt.fields.options,
				target:    tt.fields.target,
			}
			got := h.setCommand()
			sort.Strings(got.command)
			sort.Strings(tt.want.command)
			sort.Strings(got.fileNames)
			sort.Strings(tt.want.fileNames)

			if !reflect.DeepEqual(got.called, tt.want.called) {
				t.Errorf("HasuraCmd.setCommand() = %v, want.called %v", got.called, tt.want.called)
			}
			if !reflect.DeepEqual(got.target, tt.want.target) {
				t.Errorf("HasuraCmd.setCommand() = %v, want.target %v", got.target, tt.want.target)
			}
			if !reflect.DeepEqual(got.command, tt.want.command) {
				t.Errorf("HasuraCmd.setCommand() = %v, want.command %v", got.command, tt.want.command)
			}
			if !reflect.DeepEqual(got.fileNames, tt.want.fileNames) {
				t.Errorf("HasuraCmd.setCommand() = %v, want.fileNames %v", got.fileNames, tt.want.fileNames)
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
						t.Errorf("HasuraCmd.setCommand() = %v, want.options %v", got.options, tt.want.options)
					}
				}
				if len(tt.want.options) != len(got.options) {
					t.Errorf("HasuraCmd.setCommand() = %v, want.options %v", got.options, tt.want.options)
				}
			}
		})
	}
}
