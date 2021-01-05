systemctl stop datadog-custom-metrics
systemctl disable datadog-custom-metrics
rm /etc/systemd/system/datadog-custom-metrics.service
systemctl daemon-reload
systemctl reset-failed