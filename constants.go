package kivik

const (
	// KivikVersion is the version of the Kivik library.
	KivikVersion = "1.0.0-beta"
	// KivikVendor is the vendor string reported by this library.
	KivikVendor = "Kivik"
)

// SessionCookieName is the name of the CouchDB session cookie.
const SessionCookieName = "AuthSession"

// UserPrefix is the mandatory CouchDB user prefix.
// See http://docs.couchdb.org/en/2.0.0/intro/security.html#org-couchdb-user
const UserPrefix = "org.couchdb.user:"

// EndKeySuffix is a high Unicode character (0xfff0) useful for appending to an
// endkey argument, when doing a ranged search, as described here:
// http://couchdb.readthedocs.io/en/latest/ddocs/views/collation.html#string-ranges
//
// Example, to return all results with keys beginning with "foo":
//
//    rows, err := db.Query(context.TODO(), "ddoc", "view", map[string]interface{}{
//        "startkey": "foo",
//        "endkey":   "foo" + kivik.EndKeySuffix,
//    })
const EndKeySuffix = string(0xfff0)

// HTTP methods supported by CouchDB. This is almost an exact copy of the
// methods in the standard http package, with the addition of MethodCopy, and
// a few methods left out which are not used by CouchDB.
const (
	MethodGet    = "GET"
	MethodHead   = "HEAD"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
	MethodCopy   = "COPY"
)

// HTTP response codes permitted by the CouchDB API.
// See http://docs.couchdb.org/en/1.6.1/api/basics.html#http-status-codes
const (
	StatusOK                           = 200
	StatusCreated                      = 201
	StatusAccepted                     = 202
	StatusFound                        = 302
	StatusNotModified                  = 304
	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusMethodNotAllowed             = 405
	StatusRequestTimeout               = 408
	StatusConflict                     = 409
	StatusPreconditionFailed           = 412
	StatusStatusRequestEntityTooLarge  = 413
	StatusUnsupportedMediaType         = 415
	StatusRequestedRangeNotSatisfiable = 416
	StatusExpectationFailed            = 417
	StatusInternalServerError          = 500

	// StatusNotImplemented is not returned by CouchDB proper. It is used by
	// Kivik for optional features which are not implemented by some drivers.
	StatusNotImplemented = 501

	// Error status over 600 are obviously not proper HTTP errors at all. They
	// are used for kivik-generated errors of various types.

	// StatusUnknownError is used for unclassified errors generated by Kivik,
	// generally caused by a programming error.
	StatusUnknownError = 600

	// StatusNetworkError represents an error that occurred outside of the HTTP
	// transport, such as a problem contacting the remote server.  This code is
	// used to distinguish from actual HTTP-layer errors.
	StatusNetworkError = 601

	// StatusBadResponse indicates that the server responded with unrecognized
	// data.
	StatusBadResponse = 602

	// StatusIteratorUnusable indicates an improper use of an iterator, such as
	// calling an iterator prematurely, or after it was closed.
	StatusIteratorUnusable = 603
)

// Exit statuses, borrowed from Curl. Not all Curl statuses are represented here.
const (
	// Exited with an unknown failure.
	ExitUnknownFailure = 1
	// Failed to initialize.
	ExitFailedToInitialize = 2
	// URL malformed. The syntax was not correct.
	ExitStatusURLMalformed = 3
	// Write error. Kouch couldn't write data to a local filesystem or similar.
	ExitWriteError = 23

/*
5      Couldn't resolve proxy. The given proxy host could not be resolved.
6      Couldn't resolve host. The given remote host was not resolved.
7      Failed to connect to host.
8      Weird server reply. The server sent data curl couldn't parse.
18     Partial file. Only a part of the file was transferred.
22     HTTP page not retrieved. The requested url was not found or returned another error with the HTTP error code being 400 or above. This return code only appears if -f, --fail is used.
26     Read error. Various reading problems.
27     Out of memory. A memory allocation request failed.
28     Operation timeout. The specified time-out period was reached according to the conditions.
33     HTTP range error. The range "command" didn't work.
34     HTTP post error. Internal post-request generation error.
35     SSL connect error. The SSL handshaking failed.
37     FILE couldn't read file. Failed to open the file. Permissions?
43     Internal error. A function was called with a bad parameter.
45     Interface error. A specified outgoing interface could not be used.
47     Too many redirects. When following redirects, curl hit the maximum amount.
51     The peer's SSL certificate or SSH MD5 fingerprint was not OK.
52     The server didn't reply anything, which here is considered an error.
53     SSL crypto engine not found.
54     Cannot set SSL crypto engine as default.
55     Failed sending network data.
56     Failure in receiving network data.
58     Problem with the local certificate.
59     Couldn't use specified SSL cipher.
60     Peer certificate cannot be authenticated with known CA certificates.
61     Unrecognized transfer encoding.
63     Maximum file size exceeded.
65     Sending the data requires a rewind that failed.
66     Failed to initialise SSL Engine.
67     The user name, password, or similar was not accepted and curl failed to log in.
75     Character conversion failed.
76     Character conversion functions required.
77     Problem with reading the SSL CA cert (path? access rights?).
78     The resource referenced in the URL does not exist.
79     An unspecified error occurred during the SSH session.
80     Failed to shut down the SSL connection.
82     Could not load CRL file, missing or wrong format (added in 7.19.0).
83     Issuer check failed (added in 7.19.0).
85     RTSP: mismatch of CSeq numbers
86     RTSP: mismatch of Session Identifiers
89     No connection available, the session will be queued
90     SSL public key does not matched pinned public key
*/
)
