import { Catalog } from './catalog';
import { dashboardForService } from './dashboard';
import { Client } from './grafana';

const printDevelopmentDashboard = (): void => {
    const service = {
        name: 'products',
        description: 'A service related to products',
        has_http: true,
        has_grpc: true,
        github: 'http://github.com/org/products-service'
    };

    const dashboard = dashboardForService(service);
    
    console.log(JSON.stringify(dashboard.build(), null, 2));
};

const fetchServicesAndDeploy = async (): Promise<void> => {
    const grafana = Client.withConfigFromEnv(process.env);
    const catalog = Catalog.withConfigFromEnv(process.env);
    const services = await catalog.services();

    for (const service of services) {
        const dashboard = dashboardForService(service);
        const folderUid = await grafana.findOrCreateFolder(service.name);

        await grafana.persistDashboard(folderUid, dashboard);
    }

    console.log(`${services.length} dashboards deployed`);
};

(async () => {
    const deploy = process.argv.includes('--deploy');
    const help = process.argv.includes('--help') || process.argv.includes('-h');
    
    if (help) {
        console.log('Usage:');
        console.log("\t--deploy\tFetch the list of services from the catalog and deploy a dashboard for each entry");
        process.exit(1);
    }
    
    // By default, assume we're in "development mode" and print a single
    // dashboard to stdout.
    if (!deploy) {
        printDevelopmentDashboard();
    } else {
        // Otherwise, fetch the list services from the catalog and deploy a
        // dashboard for each of them
        await fetchServicesAndDeploy();
    }
})();