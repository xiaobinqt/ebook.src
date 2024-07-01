



# 帮我改一下这个正则表达式，让它能够匹配到所有的 <script type="text/markdown" data-help-name="xxx">xxx</script>

grep -iEo "<script type=\"text/markdown\" data-help-name=[\"'].+[\"']>([\s\S]*?)</script>"


 grep -iPo "<script\s*type\s*=\s*\"text/markdown\"\s*data-help-name\s*=\s*[\"'].+?[\"']>([\s\S]*?)</script>" | grep -oP ">.*?<" | grep -oP "[^><]+"


# grep -iEo "<script type=\"text/markdown\" data-help-name=[\"'].+[\"']">([\s\S]*?)<\/script>"

grep "<script type=\"text\/markdown\" data-help-name=\"linantest1008\">\n([\s\S]*?)\n<\/script>"

grep -oP '(?<=<script type="text/markdown" data-help-name="linantest1008">).*?(?=</script>)' your_file.html


---

sed -n '/<script\s*type\s*=\s*".*"\s*data-help-name\s*=\s*".*">/,/<\/script>/p' | sed '1d;$d'



cat 1.html | sed -n '/<script\s*type\s*=\s*".*"\s*data-help-name\s*=\s*".*">/{:a;N;/<\/script>/!ba;p;q;}' | sed '1d;$d'








