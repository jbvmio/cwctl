weburl="https://go.microsoft.com/fwlink/?linkid=853070"
appname="Company Portal"
app="Company Portal.app"
logandmetadir="/tmp/logandmetadir"
tempdir=$(mktemp -d)
tempfile="$tempdir/nothingHere"
log="$logandmetadir/$appname.log"

function downloadApp () {
    echo "$(date) | Starting downlading of [$appname]"
    echo "$(date) | Downloading $appname to [$tempdir]"
    cd "$tempdir"
    curl -f -s --connect-timeout 30 --retry 5 --retry-delay 60 -L -J -O "$weburl"
    if [ $? == 0 ]; then
        tempSearchPath="$tempdir/*"
        for f in $tempSearchPath; do
            tempfile=$f
        done
        echo "$(date) | Downloaded [$app] to [$tempfile]"
    else
        echo "$(date) | Failure to download [$weburl]"
    fi

}

function inspectApp () {
    echo "$(date) | Starting inspection of [$appname]"
    echo "$(date) | Contents of [$tempdir]"
    ls -lrth $tempdir
    if [[ -f "$tempfile" ]]; then
        ftype=$(file $tempfile)
        echo "$(date) | Type for [$appname] is: $ftype"
    else
        echo "$(date) | File does not Exist [$tempfile]"
    fi
}

function cleanupAll () {
    echo "$(date) | Starting cleanup of [$appname]"
    echo "$(date) | Removing [$tempdir]"
    rm -rvf $tempdir
}

function startLog() {
    if [[ ! -d "$logandmetadir" ]]; then
        echo "$(date) | Creating [$logandmetadir] to store logs"
        mkdir -p "$logandmetadir"
    fi
    exec &> >(tee -a "$log")
}

startLog

echo ""
echo "##############################################################"
echo "# $(date) | Logging start of [$appname] to [$log]"
echo "############################################################"
echo ""

downloadApp
inspectApp
cleanupAll
