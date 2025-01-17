package gotesplit

import (
	"reflect"
	"testing"
)

func TestGetTestList(t *testing.T) {
	const sample = `TestDoCreate
TestCommandGet
TestLook
TestDoGet_bulk
TestCommandList
TestCommandListUnique
TestCommandListUnknown
TestDoList_query
TestDoList_unique
TestDoList_unknownRoot
TestDoList_notPermittedRoot
TestDoList_withSystemHiddenDir
TestDoRoot
TestDetectLocalRepoRoot
TestDetectVCSAndRepoURL
TestLocalRepositoryFromFullPath
TestNewLocalRepository
TestLocalRepositoryRoots
TestList_Symlink
TestList_Symlink_In_Same_Directory
TestFindVCSBackend
TestLocalRepository_VCS
TestLocalRepositoryRoots_URLMatchLocalRepositoryRoots
TestNewRemoteRepository
TestNewRemoteRepository_vcs_error
TestNewRemoteRepository_error
TestNewURL
TestConvertGitURLHTTPToSSH
TestNewURL_err
TestFillUsernameToPath_err
TestVCSBackend
TestCvsDummyBackend
TestBranchOptionIgnoredErrors
ok      github.com/x-motemen/ghq        0.114s
TestRunInDirSilently
TestRun
TestRunInDir
TestRunSilently
ok      github.com/x-motemen/ghq/cmdutil        0.059s
?       github.com/x-motemen/ghq/hoge   [no test files]
TestLog
ok      github.com/x-motemen/ghq/logger 0.106s`

	var expect []testList = []testList{{
		pkg: "github.com/x-motemen/ghq/logger",
		list: []string{
			"TestLog",
		}}, {
		pkg: "github.com/x-motemen/ghq/cmdutil",
		list: []string{
			"TestRun",
			"TestRunInDir",
			"TestRunInDirSilently",
			"TestRunSilently",
		}}, {
		pkg: "github.com/x-motemen/ghq",
		list: []string{
			"TestBranchOptionIgnoredErrors",
			"TestCommandGet",
			"TestCommandList",
			"TestCommandListUnique",
			"TestCommandListUnknown",
			"TestConvertGitURLHTTPToSSH",
			"TestCvsDummyBackend",
			"TestDetectLocalRepoRoot",
			"TestDetectVCSAndRepoURL",
			"TestDoCreate",
			"TestDoGet_bulk",
			"TestDoList_notPermittedRoot",
			"TestDoList_query",
			"TestDoList_unique",
			"TestDoList_unknownRoot",
			"TestDoList_withSystemHiddenDir",
			"TestDoRoot",
			"TestFillUsernameToPath_err",
			"TestFindVCSBackend",
			"TestList_Symlink",
			"TestList_Symlink_In_Same_Directory",
			"TestLocalRepositoryFromFullPath",
			"TestLocalRepositoryRoots",
			"TestLocalRepositoryRoots_URLMatchLocalRepositoryRoots",
			"TestLocalRepository_VCS",
			"TestLook",
			"TestNewLocalRepository",
			"TestNewRemoteRepository",
			"TestNewRemoteRepository_error",
			"TestNewRemoteRepository_vcs_error",
			"TestNewURL",
			"TestNewURL_err",
			"TestVCSBackend",
		}}}
	got := getTestLists(sample)
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expect: %#v\ngot: %#v", expect, got)
	}
}

func TestDetectTags(t *testing.T) {
	testCases := []struct {
		input  []string
		expect string
	}{
		{[]string{"aa", "bb"}, ""},
		{[]string{"aa", "-tags", "bb"}, "-tags=bb"},
		{[]string{"aa", "--tags=ccc", "bb"}, "--tags=ccc"},
		{[]string{"aa", "-tags"}, "-tags"},
	}

	for _, tc := range testCases {
		t.Run(tc.expect, func(t *testing.T) {
			out := detectTags(tc.input)
			if out != tc.expect {
				t.Errorf("got: %s, expect: %s", out, tc.expect)
			}
		})
	}
}
