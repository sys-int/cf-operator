#!/usr/bin/env bash
kubectl delete -f ./tests/argotunnel.yaml
kubectl delete -f ./tests/argotunnelingressrule.yaml
kubectl delete -f ./tests/dnsrecord.yaml
kubectl apply -f ./tests/argotunnel.yaml
kubectl apply -f ./tests/argotunnelingressrule.yaml
kubectl apply -f ./tests/dnsrecord.yaml