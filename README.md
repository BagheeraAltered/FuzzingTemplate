# FuzzingTemplate
nuclei template for fuzzing from my endpoint list

This nuclei template checks if a list of endpoints returns a 200 response code.
It takes a list of endpoints as input and sends HTTP requests to each endpoint.
If the response code is 200 and not a webpage, it reports a finding.

nuclei does not run fuzzing templates the -itags fuzz switch needs to be included:

nuclei -l livehosts.txt -t ~/Documents/research/endpoint.yaml -itags fuzz 

Ensure the juice.txt file is the same directory as the yaml file or change the path in the yaml file
