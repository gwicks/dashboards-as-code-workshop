import { Dashboard } from "@grafana/grafana-foundation-sdk/dashboard";

export interface Manifest {
    apiVersion: string;
    kind: string;
    metadata: Record<string, string>;
    spec: any;
}

export const dashboardManifest = (folderUid: string, dashboard: Dashboard): Manifest => {
    return {
        apiVersion: 'grizzly.grafana.com/v1alpha1',
        kind: 'Dashboard',
        metadata: {
            folder: folderUid,
            name: dashboard.uid!,
        },
        spec: dashboard,
    };
};
