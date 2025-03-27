<?php

error_reporting(E_ALL);

use App\Dashboard;
use App\Grafana;
use App\Grafana\GrizzlyManifest;

require_once __DIR__.'/vendor/autoload.php';

define('MANIFESTS_DIR', './manifests');
define('DASHBOARD_FOLDER_NAME', 'Part one');

$flags = getopt('', [
    'deploy',
    'manifests',
    'help',
]);

$deploy = array_key_exists('deploy', $flags);
$manifests = array_key_exists('manifests', $flags);
$help = array_key_exists('help', $flags);

if ($help) {
    echo 'Usage:'.PHP_EOL;
    echo "\t--deploy\tGenerate and deploy the test dashboard directly to a Grafana instance".PHP_EOL;
    echo "\t--manifests\tGenerate a dashboard manifest for the test dashboard and write it to disk".PHP_EOL;
    exit(1);
}

$grafana = new Grafana\Client(Grafana\Config::fromEnv($_ENV));
$dashboard = Dashboard\Playground::create();

// Deploy the test dashboard directly to a Grafana instance.
if ($deploy) {
    $folderUid = $grafana->findOrCreateFolder(DASHBOARD_FOLDER_NAME);

    $grafana->persistDashboard($folderUid, $dashboard);

    echo "Dashboard deployed".PHP_EOL;

    exit(0);
}

// Generate a dashboard manifest for the test dashboard and write it to disk.
if ($manifests) {
    if (!is_dir(MANIFESTS_DIR)) {
        mkdir(MANIFESTS_DIR, 0777, recursive: true);
    }

    $folderUid = $grafana->findOrCreateFolder(DASHBOARD_FOLDER_NAME);

    $manifest = GrizzlyManifest::dashboard($folderUid, $dashboard);
    $filepath = MANIFESTS_DIR . DIRECTORY_SEPARATOR . $dashboard->uid . '.yaml';
    file_put_contents($filepath, $manifest->toYaml());

    echo "Manifest generated in ".MANIFESTS_DIR.PHP_EOL;

    exit(0);
}

// Assume we're in "development mode" and print the dashboard to stdout.
echo json_encode($dashboard, JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES).PHP_EOL;