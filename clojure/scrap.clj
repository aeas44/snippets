(println "yo")

(defn blank? [s] (every? #(Character/isWhitespace %) s))
(defstruct person :first-name :last-name)

; [] はリストではなくベクタという扱い。つまり、関数の引数はベクタで与えられる。
(defn hello-world [username]
    (println (format "Hello, %s" username)))
(hello-world "Yasu")
