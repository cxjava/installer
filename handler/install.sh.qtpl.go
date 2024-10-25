// Code generated by qtc from "install.sh.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line handler/install.sh.qtpl:1
package handler

//line handler/install.sh.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line handler/install.sh.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line handler/install.sh.qtpl:1
func StreamShell(qw422016 *qt422016.Writer, r Result) {
//line handler/install.sh.qtpl:1
	qw422016.N().S(`
#!/bin/bash
if [ "$DEBUG" == "1" ]; then
	set -x
fi
TMP_DIR=$(mktemp -d -t jpillora-installer-XXXXXXXXXX)
function cleanup {
	rm -rf $TMP_DIR > /dev/null
}
function fail {
	cleanup
	msg=$1
	echo "============"
	echo "Error: $msg" 1>&2
	exit 1
}
function install {
	#settings
	USER="`)
//line handler/install.sh.qtpl:19
	qw422016.E().S(r.User)
//line handler/install.sh.qtpl:19
	qw422016.N().S(`"
	PROG="`)
//line handler/install.sh.qtpl:20
	qw422016.E().S(r.Program)
//line handler/install.sh.qtpl:20
	qw422016.N().S(`"
	ASPROG="`)
//line handler/install.sh.qtpl:21
	if len(r.AsProgram) > 0 {
//line handler/install.sh.qtpl:21
		qw422016.N().S(` `)
//line handler/install.sh.qtpl:21
		qw422016.E().S(r.AsProgram)
//line handler/install.sh.qtpl:21
		qw422016.N().S(` `)
//line handler/install.sh.qtpl:21
	}
//line handler/install.sh.qtpl:21
	qw422016.N().S(`"
	MOVE="`)
//line handler/install.sh.qtpl:22
	qw422016.E().V(r.MoveToPath)
//line handler/install.sh.qtpl:22
	qw422016.N().S(`"
	RELEASE="`)
//line handler/install.sh.qtpl:23
	qw422016.E().S(r.Release)
//line handler/install.sh.qtpl:23
	qw422016.N().S(`" # `)
//line handler/install.sh.qtpl:23
	qw422016.E().S(r.ResolvedRelease)
//line handler/install.sh.qtpl:23
	qw422016.N().S(`
	INSECURE="`)
//line handler/install.sh.qtpl:24
	qw422016.E().V(r.Insecure)
//line handler/install.sh.qtpl:24
	qw422016.N().S(`"
	OUT_DIR="`)
//line handler/install.sh.qtpl:25
	if r.MoveToPath {
//line handler/install.sh.qtpl:25
		qw422016.N().S(`/usr/local/bin`)
//line handler/install.sh.qtpl:25
	} else {
//line handler/install.sh.qtpl:25
		qw422016.N().S(`$(pwd)`)
//line handler/install.sh.qtpl:25
	}
//line handler/install.sh.qtpl:25
	qw422016.N().S(`"
	GH="https://github.com"
	#bash check
	[ ! "$BASH_VERSION" ] && fail "Please use bash instead"
	[ ! -d $OUT_DIR ] && fail "output directory missing: $OUT_DIR"
	#dependency check, assume we are a standard POISX machine
	which find > /dev/null || fail "find not installed"
	which xargs > /dev/null || fail "xargs not installed"
	which sort > /dev/null || fail "sort not installed"
	which tail > /dev/null || fail "tail not installed"
	which cut > /dev/null || fail "cut not installed"
	which du > /dev/null || fail "du not installed"
	#choose an HTTP client
	GET=""
	if which curl > /dev/null; then
		GET="curl"
		if [[ $INSECURE = "true" ]]; then GET="$GET --insecure"; fi
		GET="$GET --fail -# -L"
	elif which wget > /dev/null; then
		GET="wget"
		if [[ $INSECURE = "true" ]]; then GET="$GET --no-check-certificate"; fi
		GET="$GET -qO-"
	else
		fail "neither wget/curl are installed"
	fi
	#debug HTTP
	if [ "$DEBUG" == "1" ]; then
		GET="$GET -v"
	fi
	#optional auth to install from private repos
	#NOTE: this also needs to be set on your instance of installer
	AUTH="${GITHUB_TOKEN}"
	if [ ! -z "$AUTH" ]; then
		GET="$GET -H 'Authorization: $AUTH'"
	fi
	#find OS #TODO BSDs and other posixs
	case `)
//line handler/install.sh.qtpl:25
	qw422016.N().S("`")
//line handler/install.sh.qtpl:25
	qw422016.N().S(`uname -s`)
//line handler/install.sh.qtpl:25
	qw422016.N().S("`")
//line handler/install.sh.qtpl:25
	qw422016.N().S(` in
	Darwin) OS="darwin";;
	Linux) OS="linux";;
	*) fail "unknown os: $(uname -s)";;
	esac
	#find ARCH
	if uname -m | grep -E '(arm|arch)64' > /dev/null; then
		ARCH="arm64"
		`)
//line handler/install.sh.qtpl:69
	if !r.M1Asset {
//line handler/install.sh.qtpl:69
		qw422016.N().S(`
		# no m1 assets. if on mac arm64, rosetta allows fallback to amd64
		if [[ $OS = "darwin" ]]; then
			ARCH="amd64"
		fi
		`)
//line handler/install.sh.qtpl:74
	}
//line handler/install.sh.qtpl:74
	qw422016.N().S(`
	elif uname -m | grep 64 > /dev/null; then
		ARCH="amd64"
	elif uname -m | grep arm > /dev/null; then
		ARCH="arm" #TODO armv6/v7
	elif uname -m | grep 386 > /dev/null; then
		ARCH="386"
	else
		fail "unknown arch: $(uname -m)"
	fi
	#choose from asset list
	URL=""
	FTYPE=""
	case "${OS}_${ARCH}" in`)
//line handler/install.sh.qtpl:87
	for _, n := range r.Assets {
//line handler/install.sh.qtpl:87
		qw422016.N().S(`
	"`)
//line handler/install.sh.qtpl:88
		qw422016.E().S(n.OS)
//line handler/install.sh.qtpl:88
		qw422016.N().S(`_`)
//line handler/install.sh.qtpl:88
		qw422016.E().S(n.Arch)
//line handler/install.sh.qtpl:88
		qw422016.N().S(`")
		URL="`)
//line handler/install.sh.qtpl:89
		qw422016.E().S(n.URL)
//line handler/install.sh.qtpl:89
		qw422016.N().S(`"
		FTYPE="`)
//line handler/install.sh.qtpl:90
		qw422016.E().S(n.Type)
//line handler/install.sh.qtpl:90
		qw422016.N().S(`"
		;;`)
//line handler/install.sh.qtpl:91
	}
//line handler/install.sh.qtpl:91
	qw422016.N().S(`
	*) fail "No asset for platform ${OS}-${ARCH}";;
	esac
	#got URL! download it...
	echo -n "`)
//line handler/install.sh.qtpl:95
	if r.MoveToPath {
//line handler/install.sh.qtpl:95
		qw422016.N().S(`Installing`)
//line handler/install.sh.qtpl:95
	} else {
//line handler/install.sh.qtpl:95
		qw422016.N().S(`Downloading`)
//line handler/install.sh.qtpl:95
	}
//line handler/install.sh.qtpl:95
	qw422016.N().S(`"
	echo -n " $USER/$PROG"
	if [ ! -z "$RELEASE" ]; then
		echo -n " $RELEASE"
	fi
	if [ ! -z "$ASPROG" ]; then
		echo -n " as $ASPROG"
	fi
	echo -n " (${OS}/${ARCH})"
	`)
//line handler/install.sh.qtpl:104
	if r.Search {
//line handler/install.sh.qtpl:104
		qw422016.N().S(`
	# web search, give time to cancel
	echo -n " in 5 seconds"
	for i in 1 2 3 4 5; do
		sleep 1
		echo -n "."
	done
	`)
//line handler/install.sh.qtpl:111
	} else {
//line handler/install.sh.qtpl:111
		qw422016.N().S(`
	echo "....."
	`)
//line handler/install.sh.qtpl:113
	}
//line handler/install.sh.qtpl:113
	qw422016.N().S(`
	#enter tempdir
	mkdir -p $TMP_DIR
	cd $TMP_DIR
	if [[ $FTYPE = ".gz" ]]; then
		which gzip > /dev/null || fail "gzip is not installed"
		bash -c "$GET $URL" | gzip -d - > $PROG || fail "download failed"
	elif [[ $FTYPE = ".bz2" ]]; then
		which bzip2 > /dev/null || fail "bzip2 is not installed"
		bash -c "$GET $URL" | bzip2 -d - > $PROG || fail "download failed"
	elif [[ $FTYPE = ".tar.bz" ]] || [[ $FTYPE = ".tar.bz2" ]]; then
		which tar > /dev/null || fail "tar is not installed"
		which bzip2 > /dev/null || fail "bzip2 is not installed"
		bash -c "$GET $URL" | tar jxf - || fail "download failed"
	elif [[ $FTYPE = ".tar.gz" ]] || [[ $FTYPE = ".tgz" ]]; then
		which tar > /dev/null || fail "tar is not installed"
		which gzip > /dev/null || fail "gzip is not installed"
		bash -c "$GET $URL" | tar zxf - || fail "download failed"
	elif [[ $FTYPE = ".zip" ]]; then
		which unzip > /dev/null || fail "unzip is not installed"
		bash -c "$GET $URL" > tmp.zip || fail "download failed"
		unzip -o -qq tmp.zip || fail "unzip failed"
		rm tmp.zip || fail "cleanup failed"
	elif [[ $FTYPE = ".bin" ]]; then
		bash -c "$GET $URL" > "`)
//line handler/install.sh.qtpl:137
	qw422016.E().S(r.Program)
//line handler/install.sh.qtpl:137
	qw422016.N().S(`_${OS}_${ARCH}" || fail "download failed"
	else
		fail "unknown file type: $FTYPE"
	fi
	#search subtree largest file (bin)
	TMP_BIN=$(find . -type f | xargs du | sort -n | tail -n 1 | cut -f 2)
	if [ ! -f "$TMP_BIN" ]; then
		fail "could not find find binary (largest file)"
	fi
	#ensure its larger than 1MB
	#TODO linux=elf/darwin=macho file detection?
	if [[ $(du -m $TMP_BIN | cut -f1) -lt 1 ]]; then
		fail "no binary found ($TMP_BIN is not larger than 1MB)"
	fi
	#move into PATH or cwd
	chmod +x $TMP_BIN || fail "chmod +x failed"
	DEST="$OUT_DIR/$PROG"	
	if [ ! -z "$ASPROG" ]; then
		DEST="$OUT_DIR/$ASPROG"
	fi
	#move without sudo
	OUT=$(mv $TMP_BIN $DEST 2>&1)
	STATUS=$?
	# failed and string contains "Permission denied"
	if [ $STATUS -ne 0 ]; then
		if [[ $OUT =~ "Permission denied" ]]; then
			echo "mv with sudo..."
			sudo mv $TMP_BIN $DEST || fail "sudo mv failed" 
		else
			fail "mv failed ($OUT)"
		fi
	fi
	echo "`)
//line handler/install.sh.qtpl:169
	if r.MoveToPath {
//line handler/install.sh.qtpl:169
		qw422016.N().S(`Installed at`)
//line handler/install.sh.qtpl:169
	} else {
//line handler/install.sh.qtpl:169
		qw422016.N().S(`Downloaded to`)
//line handler/install.sh.qtpl:169
	}
//line handler/install.sh.qtpl:169
	qw422016.N().S(` $DEST"
	#done
	cleanup
}
install
`)
//line handler/install.sh.qtpl:174
}

//line handler/install.sh.qtpl:174
func WriteShell(qq422016 qtio422016.Writer, r Result) {
//line handler/install.sh.qtpl:174
	qw422016 := qt422016.AcquireWriter(qq422016)
//line handler/install.sh.qtpl:174
	StreamShell(qw422016, r)
//line handler/install.sh.qtpl:174
	qt422016.ReleaseWriter(qw422016)
//line handler/install.sh.qtpl:174
}

//line handler/install.sh.qtpl:174
func Shell(r Result) string {
//line handler/install.sh.qtpl:174
	qb422016 := qt422016.AcquireByteBuffer()
//line handler/install.sh.qtpl:174
	WriteShell(qb422016, r)
//line handler/install.sh.qtpl:174
	qs422016 := string(qb422016.B)
//line handler/install.sh.qtpl:174
	qt422016.ReleaseByteBuffer(qb422016)
//line handler/install.sh.qtpl:174
	return qs422016
//line handler/install.sh.qtpl:174
}