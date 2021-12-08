# IaC

I initialized the CDK project here; however, I did not have enough time to figure out all the ins and outs of IaC over
two evenings I had to spend for this testing task. Thus, unfortunately, all of the infrastructure in AWS is manually
configured. I mentioned a couple of times that I do not posess past experiense with IaC, though I found CDK a fascinating
topic.

If I would have another week of time for this test project, I would implement the infrastructure using CDK, it
seems a fairly straighforward approach to configuration. As is, given the tight time requirement, I'm submitting the task
without this part.

I set up the infrastructure in a following manner:

1. DynamoDB table `People` was created with `first_name` and `last_name` strings as partitioning and sorting key respectively.
2. A lambda from `csv-transform` was uploaded as `724725940994.dkr.ecr.eu-west-2.amazonaws.com/csv-transformer-lambda:latest`
3. A S3 bucket `csv-container-3515e6a2` was created to serve as a host for the lambda trigger.
4. A trigger was configured from S3 to invoke the lambda on `.csv` uploads.
5. A factorial from `factorial` was uploaded as `724725940994.dkr.ecr.eu-west-2.amazonaws.com/factorial-server:latest`
6. I set up a cluster, task and a deployment via ECS using this container image. It should be currently reachable
   [here][aws-factorial-lb]. The API is dicussed in the factorial [`README.md`](../factorial/README.md).
   - As a note here: I would much prefer a K8S-based deployment approach for this, with ingress and deployment configuration
     being a part of this repository. I refer you to a repository in which I did a similar [deployment][github-rendertron]

[aws-factorial-lb]: (http://factorial-lb-866289418.eu-west-2.elb.amazonaws.com/?number=5)
[github-rendertron]: https://github.com/wizzardich/rendertron-on-k8s
