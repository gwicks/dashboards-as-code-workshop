<?php

namespace App\Grafana;

use Grafana\Foundation\Dashboard\Dashboard;
use Symfony\Component\Yaml\Yaml;

class Manifest implements \JsonSerializable
{
    public function __construct(
        public readonly string $apiVersion,
        public readonly string $kind,
        public readonly array $metadata,
        public readonly mixed $spec,
    )
    {
    }

    public static function dashboard(string $folderUid, Dashboard $dashboard): static
    {
        return new static(
            apiVersion: 'grizzly.grafana.com/v1alpha1',
            kind: 'Dashboard',
            metadata: [
                'folder' => $folderUid,
                'name' => $dashboard->uid,
            ],
            spec: $dashboard,
        );
    }

    public function jsonSerialize(): array
    {
        return [
            'apiVersion' => $this->apiVersion,
            'kind' => $this->kind,
            'metadata' => $this->metadata,
            'spec' => $this->spec,
        ];
    }

    public function toYaml(): string
    {
        $array = json_decode(json_encode($this), associative: true);
        return Yaml::dump($array, inline: 10, indent: 4, flags: Yaml::DUMP_NULL_AS_TILDE | Yaml::DUMP_EMPTY_ARRAY_AS_SEQUENCE);
    }
}