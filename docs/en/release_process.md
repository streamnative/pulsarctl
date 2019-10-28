# Release process

This guide illustrates how to perform a release for pulsarctl.

## Making the release

The steps for releasing are as follows:

1. Create the release branch
2. Promote the release
3. Write a release note

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
$ git clone git@github.com:streamnattive/pulsarctl.git
$ cd pulsarctl
$ git checkout -b branch-0.1.X origin/master
```

#### 2. Create and push the release.

```shell
# Create a tag
$ git tag -u $USER@streamnative.io v0.1.X -m 'Release v0.1.X'
# Update VERSION file
$ echo v0.1.X > VERSION

# Push both the branch and the tag to Github repo
$ git push origin branch-0.1.X
$ git push origin v0.1.X
```

#### 3. Publish a release note

Check the milestone in GitHub associated with the release. 

In the released item, add the list of the most important changes that happened in the release and a link to the associated milestone, with the complete list of all the changes. 

Update the release draft at [the release homepage of pulsarctl](https://github.com/streamnative/pulsarctl/releases)

Then publish the release draft, the binary and commands doc will auto publish to that release
