(define-module (go server)
  #:use-module (system repl server)
  #:export (socket-file
            server-start))

(define socket-file "/tmp/go-scheme.socket")

(define (server-start)
  "Spawn a UNIX domain sockert REPL in a new thread. The file is the
value of socket-file."
  (when (file-exists? socket-file)
    (delete-file socket-file))
  (spawn-server
   (make-unix-domain-server-socket #:path socket-file)))
