package tests

import (
	"path/filepath"
)

const (
	NugetRemoteRepo                        = "jfrog-cli-tests-nuget-remote-repo"
	RtServerId                             = "rtTestServerId"
	BuildAddDepsSpec                       = "simple_bad_spec.json"
	BuildAddDepsDoubleSpec                 = "double_bad_spec.json"
	BuildAddDepsBuildName                  = "cli-bad-test-build"
	NpmBuildName                           = "cli-npm-test-build"
	NugetBuildName                         = "cli-nuget-test-build"
	Out                                    = "out"
	Temp                                   = "tmp"
	DownloadSpec                           = "download_spec.json"
	BuildDownloadSpec                      = "build_download_spec.json"
	BuildDownloadSpecNoBuildNumber         = "build_download_spec_no_build_number.json"
	SimpleUploadSpec                       = "simple_upload_spec.json"
	UploadEmptyDirs                        = "upload_empty_dir_spec.json"
	DownloadEmptyDirs                      = "download_empty_dir_spec.json"
	SplitUploadSpecA                       = "split_upload_spec_a.json"
	SplitUploadSpecB                       = "split_upload_spec_b.json"
	UploadSpec                             = "upload_spec.json"
	DeleteSpec                             = "delete_spec.json"
	DeleteComplexSpec                      = "delete_complex_spec.json"
	MoveCopyDeleteSpec                     = "move_copy_delete_spec.json"
	PrepareCopy                            = "prepare_copy.json"
	Search                                 = "search.json"
	SearchGo                               = "search_go.json"
	SearchAllRepo1                         = "search_all_repo1.json"
	SearchRepo2                            = "search_repo2.json"
	SearchTxt                              = "search_txt.json"
	SearchMoveDeleteRepoSpec               = "search_move_delete_repo_spec.json"
	CopyByBuildSpec                        = "move_copy_delete_by_build_spec.json"
	CpMvDlByBuildAssertSpec                = "copy_by_build_assert_spec.json"
	GitLfsAssertSpec                       = "git_lfs_assert_spec.json"
	MoveRepositoryConfig                   = "move_repository_config.json"
	SpecsTestRepositoryConfig              = "specs_test_repository_config.json"
	VirtualRepositoryConfig                = "specs_virtual_repository_config.json"
	GitLfsTestRepositoryConfig             = "git_lfs_test_repository_config.json"
	JcenterRemoteRepositoryConfig          = "jcenter_remote_repository_config.json"
	NpmLocalRepositoryConfig               = "npm_local_repository_config.json"
	NpmRemoteRepositoryConfig              = "npm_remote_repository_config.json"
	GoLocalRepositoryConfig                = "go_local_repository_config.json"
	RepoDetailsUrl                         = "api/repositories/"
	CopyItemsSpec                          = "copy_items_spec.json"
	MavenServerIDConfig                    = "maven_server_id.yaml"
	MavenUsernamePasswordTemplate          = "maven_user_pass_template.yaml"
	GradleServerIDConfig                   = "gradle_server_id.yaml"
	GradleUseramePasswordTemplate          = "gradle_user_pass_template.yaml"
	DownloadSpecExclude                    = "download_spec_exclude.json"
	MoveCopySpecExclude                    = "move_copy_spec_exclude.json"
	DelSpecExclude                         = "delete_spec_exclude.json"
	UploadSpecExclude                      = "upload_spec_exclude.json"
	UploadSpecExcludeRegex                 = "upload_spec_exclude_regex.json"
	BuildDownloadSpecNoBuildNumberWithSort = "build_download_spec_no_build_number_with_sort.json"
	HttpsProxyEnvVar                       = "PROXY_HTTPS_PORT"
	ArchiveEntriesUpload                   = "archive_entries_upload_spec.json"
	ArchiveEntriesDownload                 = "archive_entries_download_spec.json"
)

var Repo1 = "jfrog-cli-tests-repo1"
var Repo2 = "jfrog-cli-tests-repo2"
var VirtualRepo = "jfrog-cli-tests-virtual-repo"
var JcenterRemoteRepo = "jfrog-cli-tests-jcenter-remote"
var LfsRepo = "jfrog-cli-tests-lfs-repo"
var NpmLocalRepo = "jfrog-cli-tests-npm-local-repo"
var NpmRemoteRepo = "jfrog-cli-tests-npm-remote-repo"
var GoLocalRepo = "jfrog-cli-tests-go-local-repo"

func GetTxtUploadExpectedRepo1() []string {
	return []string{
		Repo1 + "/cliTestFile.txt",
	}
}

func GetSimpleUploadExpectedRepo1() []string {
	return []string{
		Repo1 + "/flat_recursive/a3.in",
		Repo1 + "/flat_recursive/a1.in",
		Repo1 + "/flat_recursive/a2.in",
		Repo1 + "/flat_recursive/b2.in",
		Repo1 + "/flat_recursive/b3.in",
		Repo1 + "/flat_recursive/b1.in",
		Repo1 + "/flat_recursive/c2.in",
		Repo1 + "/flat_recursive/c1.in",
		Repo1 + "/flat_recursive/c3.in",
	}
}

func GetSimpleUploadSpecialCharNoRegexExpectedRepo1() []string {
	return []string{
		Repo1 + "/a1.in",
	}
}

func GetUploadSpecExcludeRepo1() []string {
	return []string{
		Repo1 + "/a1.in",
		Repo1 + "/b1.in",
		Repo1 + "/c2.in",
		Repo1 + "/c3.in",
	}
}

func GetUploadDebianExpected() []string {
	return []string{
		Repo1 + "/data/a1.in",
		Repo1 + "/data/a2.in",
		Repo1 + "/data/a3.in",
	}
}

func GetSingleFileCopy() []string {
	return []string{
		Repo2 + "/path/a1.in",
	}
}

func GetSingleFileCopyFullPath() []string {
	return []string{
		Repo2 + "/path/inner/a1.in",
	}
}

func GetSingleInnerFileCopyFullPath() []string {
	return []string{
		Repo2 + "/path/path/inner/a1.in",
	}
}

func GetFolderCopyTwice() []string {
	return []string{
		Repo2 + "/path/inner/a1.in",
		Repo2 + "/path/path/inner/a1.in",
	}
}

func GetFolderCopyIntoFolder() []string {
	return []string{
		Repo2 + "/path/path/inner/a1.in",
	}
}

func GetSingleDirectoryCopyFlat() []string {
	return []string{
		Repo2 + "/inner/a1.in",
	}
}

func GetAnyItemCopy() []string {
	return []string{
		Repo2 + "/path/inner/a1.in",
		Repo2 + "/someFile",
	}
}

func GetAnyItemCopyRecursive() []string {
	return []string{
		Repo2 + "/a/b/a1.in",
		Repo2 + "/aFile",
	}
}

func GetCopyFolderRename() []string {
	return []string{
		Repo2 + "/newPath/inner/a1.in",
	}
}

func GetAnyItemCopyUsingSpec() []string {
	return []string{
		Repo2 + "/a1.in",
	}
}

func GetExplodeUploadExpectedRepo1() []string {
	return []string{
		Repo1 + "/a/a3.in",
		Repo1 + "/a/a1.in",
		Repo1 + "/a/a2.in",
		Repo1 + "/a/b/b1.in",
		Repo1 + "/a/b/b2.in",
		Repo1 + "/a/b/b3.in",
		Repo1 + "/a/b/c/c1.in",
		Repo1 + "/a/b/c/c2.in",
		Repo1 + "/a/b/c/c3.in",
	}
}

func GetSimpleUploadExpectedRepo2() []string {
	return []string{
		Repo2 + "/flat_recursive/a3.in",
		Repo2 + "/flat_recursive/a1.in",
		Repo2 + "/flat_recursive/a2.in",
		Repo2 + "/flat_recursive/b2.in",
		Repo2 + "/flat_recursive/b3.in",
		Repo2 + "/flat_recursive/b1.in",
		Repo2 + "/flat_recursive/c2.in",
		Repo2 + "/flat_recursive/c1.in",
		Repo2 + "/flat_recursive/c3.in",
	}
}

func GetMassiveMoveExpected() []string {
	return []string{
		Repo2 + "/3_only_flat_recursive_target/a3.in",
		Repo2 + "/3_only_flat_recursive_target/b3.in",
		Repo2 + "/3_only_flat_recursive_target/c3.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a1.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a2.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a3.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b1.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b2.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b3.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c1.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c2.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c3.in",
		Repo2 + "/flat_nonrecursive_target/a1.in",
		Repo2 + "/flat_nonrecursive_target/a2.in",
		Repo2 + "/flat_nonrecursive_target/a3.in",
		Repo2 + "/flat_recursive_target/a1.in",
		Repo2 + "/flat_recursive_target/a2.in",
		Repo2 + "/flat_recursive_target/a3.in",
		Repo2 + "/flat_recursive_target/b1.in",
		Repo2 + "/flat_recursive_target/b2.in",
		Repo2 + "/flat_recursive_target/b3.in",
		Repo2 + "/flat_recursive_target/c1.in",
		Repo2 + "/flat_recursive_target/c2.in",
		Repo2 + "/flat_recursive_target/c3.in",
		Repo2 + "/no_target/a/a1.in",
		Repo2 + "/no_target/a/a2.in",
		Repo2 + "/no_target/a/a3.in",
		Repo2 + "/no_target/a/b/b1.in",
		Repo2 + "/no_target/a/b/b2.in",
		Repo2 + "/no_target/a/b/b3.in",
		Repo2 + "/no_target/a/b/c/c1.in",
		Repo2 + "/no_target/a/b/c/c2.in",
		Repo2 + "/no_target/a/b/c/c3.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a1.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a2.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a3.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a1.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a2.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a3.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b1.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b2.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b3.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c1.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c2.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c3.in",
		Repo2 + "/pattern_placeholder_target/a/a1.in",
		Repo2 + "/pattern_placeholder_target/a/a2.in",
		Repo2 + "/pattern_placeholder_target/a/a3.in",
		Repo2 + "/pattern_placeholder_target/a/b/b1.in",
		Repo2 + "/pattern_placeholder_target/a/b/b2.in",
		Repo2 + "/pattern_placeholder_target/a/b/b3.in",
		Repo2 + "/pattern_placeholder_target/a/b/c/c1.in",
		Repo2 + "/pattern_placeholder_target/a/b/c/c2.in",
		Repo2 + "/pattern_placeholder_target/a/b/c/c3.in",
		Repo2 + "/properties_target/properties_source/a/a1.in",
		Repo2 + "/properties_target/properties_source/a/a2.in",
		Repo2 + "/properties_target/properties_source/a/a3.in",
		Repo2 + "/properties_target/properties_source/a/b/b1.in",
		Repo2 + "/properties_target/properties_source/a/b/b2.in",
		Repo2 + "/properties_target/properties_source/a/b/b3.in",
		Repo2 + "/properties_target/properties_source/a/b/c/c1.in",
		Repo2 + "/properties_target/properties_source/a/b/c/c2.in",
		Repo2 + "/properties_target/properties_source/a/b/c/c3.in",
		Repo2 + "/rename_target/RENAMED.in",
		Repo2 + "/simple_placeholder_target/a/a1.in",
		Repo2 + "/simple_target/a1.in",
		Repo2 + "/flat_nonrecursive_target/b/b1.in",
		Repo2 + "/flat_nonrecursive_target/b/b2.in",
		Repo2 + "/flat_nonrecursive_target/b/b3.in",
		Repo2 + "/flat_nonrecursive_target/b/c/c1.in",
		Repo2 + "/flat_nonrecursive_target/b/c/c2.in",
		Repo2 + "/flat_nonrecursive_target/b/c/c3.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b1.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b2.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b3.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c1.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c2.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c3.in",
	}
}

func GetBuildBeforeCopyExpected() []string {
	return GetBuildBeforeMoveExpected()
}

func GetBuildCopyExpected() []string {
	return []string{
		Repo1 + "/data/a1.in",
		Repo1 + "/data/a2.in",
		Repo1 + "/data/a3.in",
		Repo1 + "/data/b1.in",
		Repo1 + "/data/b2.in",
		Repo1 + "/data/b3.in",
		Repo2 + "/data/a1.in",
		Repo2 + "/data/a2.in",
		Repo2 + "/data/a3.in",
	}
}

func GetGitLfsExpected() []string {
	return []string{
		LfsRepo + "/objects/4b/f4/4bf4c8c0fef3f5c8cf6f255d1c784377138588c0a9abe57e440bce3ccb350c2e",
	}
}

func GetBuildBeforeMoveExpected() []string {
	return []string{
		Repo1 + "/data/b1.in",
		Repo1 + "/data/b2.in",
		Repo1 + "/data/b3.in",
		Repo1 + "/data/a1.in",
		Repo1 + "/data/a2.in",
		Repo1 + "/data/a3.in",
	}
}

func GetBuildMoveExpected() []string {
	return []string{
		Repo1 + "/data/b1.in",
		Repo1 + "/data/b2.in",
		Repo1 + "/data/b3.in",
		Repo2 + "/data/a1.in",
		Repo2 + "/data/a2.in",
		Repo2 + "/data/a3.in",
	}
}

func GetBuildCopyExclude() []string {
	return []string{
		Repo1 + "/data/a1.in",
		Repo1 + "/data/a2.in",
		Repo1 + "/data/a3.in",
		Repo1 + "/data/b1.in",
		Repo1 + "/data/b2.in",
		Repo1 + "/data/b3.in",
		Repo2 + "/data/a1.in",
		Repo2 + "/data/a2.in",
		Repo2 + "/data/a3.in",
	}
}

func GetBuildDeleteExpected() []string {
	return []string{
		Repo1 + "/data/b1.in",
		Repo1 + "/data/b2.in",
		Repo1 + "/data/b3.in",
	}
}

func GetExtractedDownload() []string {
	return []string{
		filepath.Join(Out, "randFile"),
		filepath.Join(Out, "concurrent.tar.gz"),
	}
}

func GetVirtualDownloadExpected() []string {
	return []string{
		filepath.Join(Out, "a/a1.in"),
		filepath.Join(Out, "a/a2.in"),
		filepath.Join(Out, "a/a3.in"),
		filepath.Join(Out, "a/b/b1.in"),
		filepath.Join(Out, "a/b/b2.in"),
		filepath.Join(Out, "a/b/b3.in"),
		filepath.Join(Out, "a/b/c/c1.in"),
		filepath.Join(Out, "a/b/c/c2.in"),
		filepath.Join(Out, "a/b/c/c3.in"),
	}
}

func GetMassiveDownload() []string {
	return []string{
		filepath.Join(Out, "a1.in"),
		filepath.Join(Out, "a2.in"),
		filepath.Join(Out, "a3.in"),
		filepath.Join(Out, "download"),
		filepath.Join(Out, "download", "3_only_flat_recursive", "a3.in"),
		filepath.Join(Out, "download", "3_only_flat_recursive", "b3.in"),
		filepath.Join(Out, "download", "3_only_flat_recursive", "c3.in"),
		filepath.Join(Out, "download", "aql", "downloadTestResources", "a"),
		filepath.Join(Out, "download", "aql", "downloadTestResources", "a", "a1.in"),
		filepath.Join(Out, "download", "aql_flat", "a1.in"),
		filepath.Join(Out, "download", "aql_flat", "a2.in"),
		filepath.Join(Out, "download", "aql_flat", "a3.in"),
		filepath.Join(Out, "download", "aql_flat", "b1.in"),
		filepath.Join(Out, "download", "aql_flat", "b2.in"),
		filepath.Join(Out, "download", "aql_flat", "b3.in"),
		filepath.Join(Out, "download", "aql_flat", "c1.in"),
		filepath.Join(Out, "download", "aql_flat", "c2.in"),
		filepath.Join(Out, "download", "aql_flat", "c3.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "a1.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "a2.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "a3.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "b"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "b", "b1.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "b", "b2.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "b", "b3.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "b", "c"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "b", "c", "c1.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "b", "c", "c2.in"),
		filepath.Join(Out, "download", "defaults_recursive_nonflat", "downloadTestResources", "a", "b", "c", "c3.in"),
		filepath.Join(Out, "download", "flat_nonrecursive", "a1.in"),
		filepath.Join(Out, "download", "flat_nonrecursive", "a2.in"),
		filepath.Join(Out, "download", "flat_nonrecursive", "a3.in"),
		filepath.Join(Out, "download", "flat_recursive", "a1.in"),
		filepath.Join(Out, "download", "flat_recursive", "a2.in"),
		filepath.Join(Out, "download", "flat_recursive", "a3.in"),
		filepath.Join(Out, "download", "flat_recursive", "b1.in"),
		filepath.Join(Out, "download", "flat_recursive", "b2.in"),
		filepath.Join(Out, "download", "flat_recursive", "b3.in"),
		filepath.Join(Out, "download", "flat_recursive", "c1.in"),
		filepath.Join(Out, "download", "flat_recursive", "c2.in"),
		filepath.Join(Out, "download", "flat_recursive", "c3.in"),
		filepath.Join(Out, "download", "nonflat_nonrecursive", "downloadTestResources", "a", "a1.in"),
		filepath.Join(Out, "download", "nonflat_nonrecursive", "downloadTestResources", "a", "a2.in"),
		filepath.Join(Out, "download", "nonflat_nonrecursive", "downloadTestResources", "a", "a3.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "a1.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "a2.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "a3.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "b"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "b", "b1.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "b", "b2.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "b", "b3.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "b", "c", "c1.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "b", "c", "c2.in"),
		filepath.Join(Out, "download", "nonflat_recursive", "downloadTestResources", "a", "b", "c", "c3.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "a1.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "a2.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "a3.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "b", "b1.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "b", "b2.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "b", "b3.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "b", "c", "c1.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "b", "c", "c2.in"),
		filepath.Join(Out, "download", "properties", "downloadTestResources", "a", "b", "c", "c3.in"),
		filepath.Join(Out, "download", "rename", "a1.out"),
		filepath.Join(Out, "download", "simple", "a1.in"),
		filepath.Join(Out, "download", "simple_placeholder", "a"),
		filepath.Join(Out, "download", "simple_placeholder", "a", "a1.in"),
	}
}

func GetBuildDownload() []string {
	return []string{
		Out,
		filepath.Join(Out, "download"),
		filepath.Join(Out, "download", "aql_by_build"),
		filepath.Join(Out, "download", "aql_by_build", "data"),
		filepath.Join(Out, "download", "aql_by_build", "data", "a1.in"),
		filepath.Join(Out, "download", "aql_by_build", "data", "a2.in"),
		filepath.Join(Out, "download", "aql_by_build", "data", "a3.in"),
	}
}

func GetBuildDownloadDoesntExist() []string {
	return []string{
		Out,
	}
}

func GetBuildDownloadByShaAndBuild() []string {
	return []string{
		Out,
		filepath.Join(Out, "download"),
		filepath.Join(Out, "download", "aql_by_build"),
		filepath.Join(Out, "download", "aql_by_build", "data"),
		filepath.Join(Out, "download", "aql_by_build", "data", "a10.in"),
	}
}

func GetBuildDownloadByShaAndBuildName() []string {
	return []string{
		Out,
		filepath.Join(Out, "download"),
		filepath.Join(Out, "download", "aql_by_build"),
		filepath.Join(Out, "download", "aql_by_build", "data"),
		filepath.Join(Out, "download", "aql_by_build", "data", "a11.in"),
	}
}

func GetBuildSimpleDownload() []string {
	return []string{
		Out,
		filepath.Join(Out, "download"),
		filepath.Join(Out, "download", "simple_by_build"),
		filepath.Join(Out, "download", "simple_by_build", "data"),
		filepath.Join(Out, "download", "simple_by_build", "data", "b1.in"),
	}
}

func GetBuildExcludeDownload() []string {
	return []string{
		Out,
		filepath.Join(Out, "download"),
		filepath.Join(Out, "download", "aql_by_artifacts"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "a3.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "b1.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "b2.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "b3.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "c1.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "c3.in"),
	}
}

func GetBuildExcludeDownloadBySpec() []string {
	return []string{
		Out,
		filepath.Join(Out, "download"),
		filepath.Join(Out, "download", "aql_by_artifacts"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "a2.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "b1.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "b2.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "b3.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "c1.in"),
		filepath.Join(Out, "download", "aql_by_artifacts", "data", "c3.in"),
	}
}

func GetMassiveUpload() []string {
	return []string{
		Repo1 + "/spec-copy-test/3_only_flat_recursive/a3.in",
		Repo1 + "/spec-copy-test/3_only_flat_recursive/b3.in",
		Repo1 + "/spec-copy-test/3_only_flat_recursive/c3.in",
		Repo1 + "/spec-copy-test/copy-to-existing/a1.in",
		Repo1 + "/spec-copy-test/copy-to-existing/a2.in",
		Repo1 + "/spec-copy-test/copy-to-existing/a3.in",
		Repo1 + "/spec-copy-test/copy-to-existing/b1.in",
		Repo1 + "/spec-copy-test/copy-to-existing/b2.in",
		Repo1 + "/spec-copy-test/copy-to-existing/b3.in",
		Repo1 + "/spec-copy-test/copy-to-existing/c1.in",
		Repo1 + "/spec-copy-test/copy-to-existing/c2.in",
		Repo1 + "/spec-copy-test/copy-to-existing/c3.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/a1.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/a2.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/a3.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/b1.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/b2.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/b3.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/c1.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/c2.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/c3.in",
		Repo1 + "/spec-copy-test/flat_nonrecursive/a1.in",
		Repo1 + "/spec-copy-test/flat_nonrecursive/a2.in",
		Repo1 + "/spec-copy-test/flat_nonrecursive/a3.in",
		Repo1 + "/spec-copy-test/flat_recursive/a1.in",
		Repo1 + "/spec-copy-test/flat_recursive/a2.in",
		Repo1 + "/spec-copy-test/flat_recursive/a3.in",
		Repo1 + "/spec-copy-test/flat_recursive/b1.in",
		Repo1 + "/spec-copy-test/flat_recursive/b2.in",
		Repo1 + "/spec-copy-test/flat_recursive/b3.in",
		Repo1 + "/spec-copy-test/flat_recursive/c1.in",
		Repo1 + "/spec-copy-test/flat_recursive/c2.in",
		Repo1 + "/spec-copy-test/flat_recursive/c3.in",
		Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a1.in",
		Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a2.in",
		Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a3.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a1.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a2.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a3.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b1.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b2.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b3.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c1.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c2.in",
		Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c3.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/a1.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/a2.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/a3.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/b1.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/b2.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/b3.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c1.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c2.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c3.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a$+~&^a#/a1.in",
		Repo1 + "/spec-copy-test/properties/testsdata/c#/a#1.in",
		Repo1 + "/spec-copy-test/defaults_recursive_nonflat/a#1.in",
		Repo1 + "/spec-copy-test/copy-to-existing/a#1.in",
		Repo1 + "/spec-copy-test/simple/a1.in",
	}
}

func GetPropsExpected() []string {
	return []string{
		Repo1 + "/spec-copy-test/properties/testsdata/a/a1.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/b1.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/a2.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/b2.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c1.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/a3.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/b3.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c2.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c3.in",
		Repo1 + "/spec-copy-test/properties/testsdata/a$+~&^a#/a1.in",
		Repo1 + "/spec-copy-test/properties/testsdata/c#/a#1.in",
	}
}

func GetDelete1() []string {
	return []string{
		Repo2 + "/3_only_flat_recursive_target/a3.in",
		Repo2 + "/3_only_flat_recursive_target/b3.in",
		Repo2 + "/3_only_flat_recursive_target/c3.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a1.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a2.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a3.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b1.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b2.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b3.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c1.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c2.in",
		Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c3.in",
		Repo2 + "/flat_nonrecursive_target/a1.in",
		Repo2 + "/flat_nonrecursive_target/a2.in",
		Repo2 + "/flat_nonrecursive_target/a3.in",
		Repo2 + "/flat_recursive_target/a1.in",
		Repo2 + "/flat_recursive_target/a2.in",
		Repo2 + "/flat_recursive_target/a3.in",
		Repo2 + "/flat_recursive_target/b1.in",
		Repo2 + "/flat_recursive_target/b2.in",
		Repo2 + "/flat_recursive_target/b3.in",
		Repo2 + "/flat_recursive_target/c1.in",
		Repo2 + "/flat_recursive_target/c2.in",
		Repo2 + "/flat_recursive_target/c3.in",
		Repo2 + "/no_target/a/a1.in",
		Repo2 + "/no_target/a/a2.in",
		Repo2 + "/no_target/a/a3.in",
		Repo2 + "/no_target/a/b/b1.in",
		Repo2 + "/no_target/a/b/b2.in",
		Repo2 + "/no_target/a/b/b3.in",
		Repo2 + "/no_target/a/b/c/c1.in",
		Repo2 + "/no_target/a/b/c/c2.in",
		Repo2 + "/no_target/a/b/c/c3.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a1.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a2.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a3.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a1.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a2.in",
		Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a3.in",
		Repo2 + "/pattern_placeholder_target/a/a1.in",
		Repo2 + "/pattern_placeholder_target/a/a2.in",
		Repo2 + "/pattern_placeholder_target/a/a3.in",
		Repo2 + "/pattern_placeholder_target/a/b/b1.in",
		Repo2 + "/pattern_placeholder_target/a/b/b2.in",
		Repo2 + "/pattern_placeholder_target/a/b/b3.in",
		Repo2 + "/pattern_placeholder_target/a/b/c/c1.in",
		Repo2 + "/pattern_placeholder_target/a/b/c/c2.in",
		Repo2 + "/pattern_placeholder_target/a/b/c/c3.in",
		Repo2 + "/properties_target/properties_source/a/a1.in",
		Repo2 + "/properties_target/properties_source/a/a2.in",
		Repo2 + "/properties_target/properties_source/a/a3.in",
		Repo2 + "/properties_target/properties_source/a/b/b1.in",
		Repo2 + "/properties_target/properties_source/a/b/b2.in",
		Repo2 + "/properties_target/properties_source/a/b/b3.in",
		Repo2 + "/properties_target/properties_source/a/b/c/c1.in",
		Repo2 + "/properties_target/properties_source/a/b/c/c2.in",
		Repo2 + "/properties_target/properties_source/a/b/c/c3.in",
		Repo2 + "/rename_target/RENAMED.in",
		Repo2 + "/simple_placeholder_target/a/a1.in",
		Repo2 + "/simple_target/a1.in",
		Repo2 + "/flat_nonrecursive_target/b/b1.in",
		Repo2 + "/flat_nonrecursive_target/b/b2.in",
		Repo2 + "/flat_nonrecursive_target/b/b3.in",
		Repo2 + "/flat_nonrecursive_target/b/c/c1.in",
		Repo2 + "/flat_nonrecursive_target/b/c/c2.in",
		Repo2 + "/flat_nonrecursive_target/b/c/c3.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b1.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b2.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b3.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c1.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c2.in",
		Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c3.in",
	}
}

func GetDeleteDisplyedFiles() []string {
	return []string{
		Repo2 + "/3_only_flat_recursive_source/a/b/",
		Repo2 + "/3_only_flat_recursive_source/a/a1.in",
		Repo2 + "/3_only_flat_recursive_source/a/a2.in",
		Repo2 + "/3_only_flat_recursive_source/a/a3.in",
		Repo2 + "/flat_recursive_source/a/b/c/",
		Repo2 + "/flat_recursive_source/a/b/b1.in",
		Repo2 + "/flat_recursive_source/a/b/b2.in",
		Repo2 + "/flat_recursive_source/a/b/b3.in",
		Repo2 + "/defaults_recursive_nonflat_source/a/a1.in",
		Repo2 + "/defaults_recursive_nonflat_source/a/a2.in",
		Repo2 + "/defaults_recursive_nonflat_source/a/a3.in",
		Repo2 + "/defaults_recursive_nonflat_source/a/b/",
		Repo2 + "/flat_nonrecursive_source/a/b/c/",
		Repo2 + "/flat_nonrecursive_source/a/b/b1.in",
		Repo2 + "/flat_nonrecursive_source/a/b/b2.in",
		Repo2 + "/flat_nonrecursive_source/a/b/b3.in",
	}
}

func GetMavenDeployedArtifacts() []string {
	return []string{
		Repo1 + "/org/jfrog/cli-test/1.0/cli-test-1.0.jar",
		Repo1 + "/org/jfrog/cli-test/1.0/cli-test-1.0.pom",
	}
}

func GetGradleDeployedArtifacts() []string {
	return []string{
		Repo1 + "/minimal-example/1.0/minimal-example-1.0.jar",
	}
}

func GetNpmDeployedScopedArtifacts() []string {
	return []string{
		NpmLocalRepo + "/@jscope/jfrog-cli-tests/-/jfrog-cli-tests-1.0.0.tgz",
	}
}
func GetNpmDeployedArtifacts() []string {
	return []string{
		NpmLocalRepo + "/jfrog-cli-tests/-/jfrog-cli-tests-1.0.0.tgz",
	}
}

func GetSortAndLimit() []string {
	return []string{
		Out,
		filepath.Join(Out, "download"),
		filepath.Join(Out, "download", "sort_limit"),
		filepath.Join(Out, "download", "sort_limit", "data"),
		filepath.Join(Out, "download", "sort_limit", "data", "a1.in"),
		filepath.Join(Out, "download", "sort_limit", "data", "b"),
		filepath.Join(Out, "download", "sort_limit", "data", "b", "c"),
		filepath.Join(Out, "download", "sort_limit", "data", "b", "c", "c1.in"),
		filepath.Join(Out, "download", "sort_limit", "data", "b", "c", "c2.in"),
		filepath.Join(Out, "download", "sort_limit", "data", "b", "c", "c3.in"),
	}
}

func GetBuildDownloadByShaAndBuildNameWithSort() []string {
	return []string{
		filepath.Join(Out, "download", "sort_limit_by_build"),
		filepath.Join(Out, "download", "sort_limit_by_build", "data"),
		filepath.Join(Out, "download", "sort_limit_by_build", "data", "a11.in"),
	}
}

func GetBuildArchiveEntriesDownloadCli() []string {
	return []string{
		Out,
		filepath.Join(Out, "a.zip"),
		filepath.Join(Out, "b.zip"),
		filepath.Join(Out, "c.zip"),
	}
}

func GetBuildArchiveEntriesSpecificPathDownload() []string {
	return []string{
		Out,
		filepath.Join(Out, "b.zip"),
	}
}

func GetBuildArchiveEntriesDownloadSpec() []string {
	return []string{
		Out,
		filepath.Join(Out, "d.zip"),
	}
}
