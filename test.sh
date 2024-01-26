sanity-check-oura-webhook() {
tempfile=`mktemp`
 tempfile2=`mktemp`
 challenge="this is a challenge!"
 raw_url="$1"
 # full_url="${1}?verification_token=this_is_a_verification_token&challenge=${challenge}"
 echo "raw_url: $raw_url"
 challenge="this is a challenge!"
 token="this is a token!"
 curl --verbose --url-query "verification_token=${token}" --url-query "challenge=${challenge}" "${raw_url}" > $tempfile 2>$tempfile2
 response_challenge=$(cat "$tempfile" | jq -r .challenge)
 echo "response_challenge: $response_challenge"
 if [ "$response_challenge" != "$challenge" ]
 then
   echo "challenge ($challenge) != response_challenge ($response_challenge)" >&2
   challenge_md5="$(echo "$challenge" | $md5)"
   response_challenge_md5="$(echo "$response_challenge" | $md5)"
    echo "challenge_md5: $challenge_md5" >&2
    echo "response_challenge_md5: $response_challenge_md5" >&2
    echo "tempfile: $tempfile" >&2
    echo "tempfile2: $tempfile2" >&2
    return 2
  fi
  return 0
}


 export url=localhost:3001
 echo checking "$url"
 sanity-check-oura-webhook "$url" || echo "failed check on url: $url"
