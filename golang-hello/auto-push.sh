git add .
read -p "Enter commit message (default: 'auto push'): " commit_message
commit_message=${commit_message:-auto push}
git commit -m "$commit_message"
git push