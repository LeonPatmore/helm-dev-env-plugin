echo "Starting simple install script."
version="$(cat plugin.yaml | grep "version" | cut -d ' ' -f 2)"

echo "Installing version $version"

go build -o $HELM_PLUGIN_DIR/bin/helm-plugin
chmod +x $HELM_PLUGIN_DIR/bin/helm-plugin
