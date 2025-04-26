package lab;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.dashboard.*;
import lab.grafana.Manifest;

public class Main {
    public static void main(String[] args) throws JsonProcessingException {
        Dashboard dashboard = Playground.dashboard().build();
        com.grafana.foundation.resource.Manifest manifest = Manifest.dashboard("", dashboard);

        System.out.println(manifest.toJSON());
    }
}
