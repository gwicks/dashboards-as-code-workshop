import { Dashboard } from "@grafana/grafana-foundation-sdk/dashboard";

export interface Manifest {
    apiVersion: string;
    kind: string;
    metadata: Record<string, any>;
    spec: any;
}

export const dashboardManifest = (folderUid: string, dashboard: Dashboard): Manifest => {
    return {
        apiVersion: 'dashboard.grafana.app/v1alpha1',
        kind: 'Dashboard',
        metadata: {
            annotations: {
                'grafana.app/folder': folderUid,
            },
            name: dashboard.uid!,
        },
        spec: dashboard,
    };
};
