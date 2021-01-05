
cat <<EOF
datadog-custom-metrics has been installed as a systemd service.
To start/stop datadog-custom-metrics:
sudo systemctl start datadog-custom-metrics
sudo systemctl stop datadog-custom-metrics
To enable/disable datadog-custom-metrics starting automatically on boot:
sudo systemctl enable datadog-custom-metrics
sudo systemctl disable datadog-custom-metrics
To reload datadog-custom-metrics:
sudo systemctl restart datadog-custom-metrics
To view datadog-custom-metrics logs:
journalctl -f -u datadog-custom-metrics
EOF