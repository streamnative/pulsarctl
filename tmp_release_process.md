# Release process

This guide illustrates how to perform a release for pulsarctl.

## Making the release

The steps for releasing are as follows:

1. Create the release branch
2. Update package version and tag
3. Build and inspect the artifacts
4. Write a release note
5. Promote the release
6. Update release notes

## Steps in detail

#### 1. Create the release branch

We are going to create a branch from `master` to `branch-v0.1.X`
where the tag will be generated and where new fixes will be
applied as part of the maintenance for the release.

The branch needs only to be created when creating major releases,
and not for patch releases.

Eg: When creating `v0.1.1` release, will be creating
the branch `branch-0.1.1`, but for `v0.1.2` we
would keep using the old `branch-0.1.1`.

In these instructions, I'm referring to an fictitious release `0.1.X`. Change the release version in the examples
accordingly with the real version.

It is recommended to create a fresh clone of the repository to avoid any local files to interfere in the process:

```shell
git clone git@github.com:streamnattive/pulsarctl.git
cd pulsarctl
git checkout -b branch-0.1.X origin/master
```

#### 2. Update package version and tag

During the release process, you can  create a "candidate" tag which will get promoted to the "real" final tag after verification and approval.

```
# Bump to the release version

go build -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=pulsarctl-V0.1.X" .

# Commit 
git add .
git commit -m "Release 0.1.X" -a

# Create a "candidate" tag
git tag -u $USER@apache.org v0.1.X-candidate-1 -m 'Release v0.1.X-candidate-1'

# Push both the branch and the tag to Github repo
git push origin branch-0.1.X
git push origin v0.1.X-candidate-1
```

#### 3. Build and inspect an artifact.

```
go build -o pulsarctl main.go
```

After the build, there will be generated `pulsarctl` file.

#### 4. Write a release note

Check the milestone in GitHub associated with the release. 

In the released item, add the list of the most important changes that happened in the release and a link to the associated milestone, with the complete list of all the changes. 

#### 5. Promote the release.

```
$ git checkout branch-0.1.X
$ git tag -u $USER@apache.org v0.1.X -m 'Release v0.1.X'
$ git push origin v0.1.X
```

Publish the package to the GitHub release repo.

```
$ mkdir -p temp-release
$ cd temp-release
```

#### 6. Update the release note

Add the release notes to [the release homepage of pulsarctl](https://github.com/streamnative/pulsarctl/releases)
