FROM scratch

# replace mobydick.txt with your log file
ADD mobydick.txt searches.txt
ADD server server

ENV QUERY_FILE searches.txt
# or instead of file, use a url
#ENV QUERY_URL <url with log data>

ENV PORT 80
EXPOSE 80

ENTRYPOINT ["/server"]