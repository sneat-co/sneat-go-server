export GCLOUD_PROJECT="demo-local-sneat-app"
export FIREBASE_AUTH_EMULATOR_HOST="localhost:9099"
export FIRESTORE_EMULATOR_HOST="localhost:8080"
export SNEAT_TG_DEV_BOTS="AlextDevBot:sneat_bot"
#export GOOGLE_APPLICATION_CREDENTIALS="/Users/alexandertrakhimenok/projects/sneat/sneat-team-go/private_keys/demo-sneat.json"

SCRIPT_DIR=$(dirname "$(readlink -f "$0" || realpath "$0")")

echo "Started serve_gae.sh:, SCRIPT_DIR=$SCRIPT_DIR"
cd "$SCRIPT_DIR" || exit

echo "Running 'go mod tidy'..."
go mod tidy

echo "Starting local GAE server..."
go run .
