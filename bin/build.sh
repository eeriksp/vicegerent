# Run from the project root directory
# TODO not tested

# remove any old snap packages
rm *.snap

# build the snap
snapcraft

# the name of the produces snap will depend on the version, rename it to something specific
mv *.snap vicegerent.snap

# copy the snap to the Debian package
cp vicegerent.snap ./deb/usr/local/vicegerent.snap

# build the Debian package
dpkg-deb --build deb
mv deb.deb vicegerent.deb

# move the artifacts to the releases directory
mv vicegerent.snap ./releases/vicegerent.snap
mv vicegerent.deb ./releases/vicegerent.deb
