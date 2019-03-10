(define-module (go eval)
  #:export (go-eval))

(define (go-eval x)
  (let* ((result #nil) (error #nil))
    (catch #t
      (lambda ()
        (set! result (eval-string x)))
      (lambda (key . parameters)
        (set! error (format #f "Uncaught throw to '~a: ~a\n" key parameters))))
    (list result error)))
