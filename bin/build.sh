# Run from the project root directory with the version number as the first argument

set -eu
echo "Starting to build Vicegerent v$1"

# remove any old snap packages
echo "Remove old snaps"
rm *.snap || true # `true` prevents the error is no `*.snap` files are present

# build the snap (the name of the produces snap will depend on the version, rename it to something specific)
echo "Build the snap"
snapcraft
mv *.snap vicegerent.snap

# build the Debian package
echo "Build the .deb package"
cp vicegerent.snap ./deb/usr/local/vicegerent.snap
dpkg-deb --build deb
mv deb.deb vicegerent.deb

# move the artifacts to the releases directory
echo "Move the artifacts to the releases directory"
mv vicegerent.snap docs/releases/snap/vicegerent.snap
mv vicegerent.deb docs/releases/debian/vicegerent-"$1"_all.deb

# prepare the debian repository
echo "Prepare the Debian repository"
cd docs/releases/debian
dpkg-scanpackages . | gzip -c9  > Packages.gz