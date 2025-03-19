import { dashboardForService } from './dashboard';

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
    const service = {
        name: 'products',
        description: 'A service related to products',
        has_http: true,
        has_grpc: true,
        github: 'http://github.com/org/products-service'
    };

    const dashboard = dashboardForService(service);
    
    console.log(JSON.stringify(dashboard.build(), null, 2));

    process.exit(0);
}