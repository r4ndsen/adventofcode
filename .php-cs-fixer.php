<?php

$rules = [
    '@Symfony'                    => true,
    'concat_space'                => ['spacing' => 'one'],
    'no_superfluous_phpdoc_tags'  => [
        'remove_inheritdoc' => true,
    ],
    'yoda_style'                  => [ // do not change existing yoda style
        'equal'     => null,
        'identical' => null,
    ],
    'phpdoc_summary'              => false,
    'phpdoc_line_span'            => [
        'method'   => 'single',
        'property' => 'single',
    ],
    'global_namespace_import'     => ['import_classes' => true],
    'phpdoc_to_comment'           => false,
    'class_attributes_separation' => [
        'elements' => [
            'method'   => 'one',
            'property' => 'none',
        ],
    ],
    'single_line_comment_style'   => true,
    'phpdoc_tag_type'             => false,
    'binary_operator_spaces'      => [
        'operators' => [
            // '=>' => 'align_single_space_minimal_by_scope',
            '=>' => 'align_single_space',
            '='  => 'align_single_space_minimal',
        ],
    ],
    'php_unit_method_casing'      => false,
    'ordered_class_elements'      => [
        'order' => [
            'use_trait',
            'case',
            'constant_public',
            'constant_protected',
            'constant_private',
            'property_public',
            'property_protected',
            'property_private',
            'construct',
            'destruct',
            'magic',
            'phpunit',
            'method_public',
            'method_protected',
            'method_private',
        ],
    ],
];

$config = new PhpCsFixer\Config();
$config->setRules($rules);

return $config;
