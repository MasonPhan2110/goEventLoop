commit=$1
replace=" "
git add . && git commit -m "${commit//[-]/$replace}" && git push origin $branch