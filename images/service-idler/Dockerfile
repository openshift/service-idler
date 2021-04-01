FROM registry.ci.openshift.org/ocp/4.8:base

RUN INSTALL_PKGS="origin-service-idler" && \
    yum --enablerepo=origin-local-release install -y ${INSTALL_PKGS} && \
    rpm -V ${INSTALL_PKGS} && \
    yum clean all

LABEL io.k8s.display-name="OpenShift Container Platform Service Idler" \
      io.k8s.description="This is a component of OpenShift Container Platform which controls the idling and unidling scalable resources" \
      io.openshift.tags="openshift"

# the service-idler doesn't require root
USER 1001

ENTRYPOINT ["/usr/bin/service-idler", "--logtostderr", "--install-crds=false"]
