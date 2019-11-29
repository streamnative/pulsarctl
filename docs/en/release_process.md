# Release process

This guide illustrates how to perform a release for pulsarctl.

## Making the release

The steps for releasing are as follows:

1. Prepare for releasing
2. Create the release branch
3. Promote the release
4. Write a release note
5. Check the release
6. Announce the release

## Steps in detail

#### 1. Prepare for releasing

Create a new milestone and move those pull requests that can not
publish in this release to the new milestone.

Update the VERSION file and stable.txt file to the new version.
And send a pull request for updating the files.

#### 2. Create the release branch

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

#### 3. Create and push the release

```shell
# Create a tag
$ git tag -u $USER@streamnative.io v0.1.X -m 'Release v0.1.X'

# Push both the branch and the tag to Github repo
$ git push origin branch-0.1.X
$ git push origin v0.1.X
```

#### 4. Publish a release note

Check the milestone in GitHub associated with the release. 

In the released item, add the list of the most important changes that happened in the release and a link to the associated milestone, with the complete list of all the changes. 

Update the release draft at [the release homepage of pulsarctl](https://github.com/streamnative/pulsarctl/releases)

Then publish the release draft, the binary and commands doc will auto publish to that release

#### 5. Check the release

View the doc site https://streamnative.io/docs/pulsarctl/vx.y.z/ to
check the docs have been published successfully.

Use the `install.sh` to install the Pulsarctl and make sure the
version of the download Pulsarctl is right.

Add the latest version to the `Available Release` in `README.md`.

Delete the previous milestone.

#### 6. Announce the release

After publishing and checking the release. Please work
with Growth to announce we are release the new version
of the Pulsarctl.
