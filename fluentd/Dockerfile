# FROM fluent/fluentd:v1.16.7-debian-1.1

# USER root


# RUN ["gem", "install", "fluent-plugin-elasticsearch", "--no-document"]

# # Copy configuration files
# COPY ./conf/fluent.conf /fluentd/etc/

# # Copy entrypoint script
# COPY entrypoint.sh /bin/
# RUN chmod +x /bin/entrypoint.sh

# # Switch back to fluent user
# USER fluent

# # Set entrypoint
# ENTRYPOINT ["/bin/entrypoint.sh"]
# # Install dependencies
# RUN apt-get update && apt-get install -y sudo ruby-dev build-essential \
#     && apt-get clean

# # Reinstall elasticsearch plugin for kibana (if any)
# RUN gem uninstall elasticsearch || true \
#     && gem install elasticsearch -v 7.11.0 --no-document \
#     && sudo gem install fluent-plugin-elasticsearch  --no-document \
#     && sudo gem sources --clear-all

FROM fluent/fluentd:v1.16.7-debian-1.1

USER root

# Install Fluentd plugins
RUN gem install fluent-plugin-elasticsearch --no-document

# Copy Fluentd configuration
COPY conf/fluent.conf /fluentd/etc/

# Set correct permissions
RUN chmod 644 /fluentd/etc/fluent.conf

# Switch back to the fluent user
USER fluent

# Command to run Fluentd
CMD ["fluentd", "-c", "/fluentd/etc/fluent.conf", "-v"]
