# TweetOnAction
Send Tweet on Twitter on GithubAction triggers

# Quick Start
Please create a PayPal developer account to use this application. https://developer.twitter.com/en/apply-for-access

Once account is created, you will need to have to create an App and get it's _Consumer API Keys_ and _Access token & access token secret_

# Usage
```yaml
- name: Retag to latest
  uses: Z-M-Huang/TweetOnAction@v1
  with:
    consumer_key: xvz1evFS4wEEPTGEFPHBog
    consumer_secret: kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw
    access_token: 370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb
    token_secret: LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE
    content: First Twee on GithubAction!
```