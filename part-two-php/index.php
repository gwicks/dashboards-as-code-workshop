<?php

use App\Catalog;
use App\Dashboard;
use App\Grafana;

require_once __DIR__.'/vendor/autoload.php';

$flags = getopt('', [
    'deploy',
    'help',
]);

$deploy = array_key_exists('deploy', $flags);
$help = array_key_exists('help', $flags);

if ($help) {
    echo 'Usage:'.PHP_EOL;
    echo "\t--deploy\tFetch the list of services from the catalog and deploy a dashboard for each entry".PHP_EOL;
    exit(1);
}

// By default, assume we're in "development mode" and print a single
// dashboard to stdout.
if (!$deploy) {
    $service = new Catalog\Service(
        'products',
        'A service related to products',
        true,
        true,
        'http://github.com/org/products-service'
    );

    $dashboard = Dashboard\Overview::forService($service);
    
    echo json_encode($dashboard, JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES).PHP_EOL;

    exit(0);
}

// Otherwise, fetch the list services from the catalog and deploy a
// dashboard for each of them
$grafana = new Grafana\Client(Grafana\Config::fromEnv($_ENV));
$catalog = new Catalog\Client(Catalog\Config::fromEnv($_ENV));
$services = $catalog->services();

foreach ($services as $service) {
    $dashboard = Dashboard\Overview::forService($service);
    $folderUid = $grafana->findOrCreateFolder($service->name);

    $grafana->persistDashboard($folderUid, $dashboard);
}

$servicesCount = count($services);
echo "{$servicesCount} dashboards deployed".PHP_EOL;