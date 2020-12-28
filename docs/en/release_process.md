# Release process

This guide illustrates how to perform a release for pulsarctl.

# Naming convention

All StreamNative repository should following the naming convention:

- Branch name: `branch-X.Y.Z`
- Tag name: `vX.Y.Z(.M)`(stable)
- Release candidate tag: `vX.Y.Z(.M)-rc-$(date +%Y%m%d%H%M)` (unstable)

`(.M)`  means our internal version release number, most of our repository is an extensions/tools for apache/pulsar. To keep track on the repository is produced by which version of the apache/pulsar, we will carry the apache/pulsar version number and using the `.M` to represent our internal release version. And all repository should keep align with streamnative/pulsar.

There has two type of the tags, one is stable `vX.Y.Z(.M)`, and another is unstable `vX.Y.Z(.M)-rc-$(date +%y%m%d%H%M)`. A stable tag represent this release is a verified release, and an unstable tag represent this release is not verified.

## Release workflow

The steps for releasing are as follows:

1. Prepare for a release
2. Create the release branch
3. Push the release
4. Write a release note
5. Check the release
6. Update the homebrew formula
7. Announce the release

## Steps in detail

1. Prepare for a release

Create a new milestone and move the pull requests that can not
be published in this release to the new milestone.

Update the information of the new release to the `VERSION` file
and `stable.txt` file and send a PR for requesting the changes.

2. Create the release branch

We are going to create a branch from `master` to `branch-vx.y.z`
where the tag will be generated and where new fixes will be
applied as part of the maintenance for the release. `x.y.z`
is the version of the release.

The branch needs only to be created when creating major releases,
and not for patch releases.

Eg: When creating `v0.1.1` release, will be creating
the branch `branch-0.1.1`, but for `v0.1.2` we
would keep using the old `branch-0.1.1`.

In these instructions, I'm referring to an fictitious release `x.y.z`.
Change the release version in the examples accordingly with the real version.

It is recommended to create a fresh clone of the repository to 
avoid any local files to interfere in the process:

```shell
$ git clone git@github.com:streamnattive/pulsarctl.git
$ cd pulsarctl

# Create a branch
$ git checkout -b branch-x.y.z origin/master

# Create a tag
$ git tag -u $USER@streamnative.io vx.y.z -m 'Release vx.y.z'
```

3. Push the release

```shell
# Push both the branch and the tag to Github repo
$ git push origin branch-x.y.z
$ git push origin vx.y.z
```

4. Publish a release note

Check the milestone in GitHub associated with the release. 

In the released item, add the list of the most important changes 
that happened in the release and a link to the associated milestone,
with the complete list of all the changes. 

Update the release draft at [the release homepage of pulsarctl](https://github.com/streamnative/pulsarctl/releases)

Then the release draft, binary, and command doc will be published
 o that release automatically.

5. Check the release

(1) Visit the Pulsarctl website https://docs.streamnative.io/pulsarctl/vx.y.z/ to
check the docs have been published successfully. `vx.y.z` is the version of the
release. For example, https://streamnative.io/docs/publisctl/v0.1.0/.

(2) Use the install command `sh -c "$(curl -fsSL https://raw.githubusercontent.com/streamnative/pulsarctl/master/install.sh)"`
to install the Pulsarctl and make sure the version of the downloaded Pulsarctl is right.

(3) Add the latest version to the `Available Release` in `README.md`.

(4) Close the previous milestone.

6. Update the homebrew formula

Create a pull request to the [homebrew-streamnatve](https://github.com/streamnative/homebrew-streamnative) 
for updating the `tag` and `revision` in the 
[pulsarctl formula](https://github.com/streamnative/homebrew-streamnative/blob/master/Formula/pulsarctl.rb).

7. Announce the release

After publishing and checking the release, work with Growth team
to announce that a new version of Pulsarctl is released.
