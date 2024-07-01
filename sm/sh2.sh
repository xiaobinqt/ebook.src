


tar xf /tmp/1686708475833825538.tgz package/linantest1008.html -O | sed -n '/<script type="(text\/markdown|text\/html)" data-help-name=".*">/,/<\/script>/p' | sed '1d;$d'

