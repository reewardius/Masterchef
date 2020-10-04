module.exports = function(grunt) {

    grunt.initConfig({
        concat: {
            options: {
                stripBanners: {
                    block: true,
                    line: true
                },
            },
            css: {
                src: [
                    'node_modules/bootstrap/dist/css/bootstrap.min.css',
                    'build/style.min.css',
                ],
                dest: 'build/all.min.css'
            },
            js: {
                src: [
                    'node_modules/sortablejs/dist/sortable.umd.js',
                    'src/js/main.js',
                ],
                dest: 'build/all.js',
            },
        },
        htmlbuild: {
            compile: {
                options: {
                    styles: { css: 'build/all.min.css' },
                    scripts: { js: 'build/all.min.js' },
                },
                src: 'build/index.html',
                dest: 'build/'
            }
        },
        htmlmin: {
            compile: {
                options: {
                    removeComments: true,
                    collapseWhitespace: true
                },
                files: { 'build/index.html': 'build/index.html' }
            }
        },
        pug: {
            compile: {
                options: { pretty: true },
                files: { 'build/index.html': 'src/index.pug' }
            }
        },
        sass: {
            compile: {
                options: {
                    style: 'compressed',
                    noSourceMap: true,
                },
                files: { 'build/style.min.css': 'src/sass/style.sass' }
            }
        },
        uglify: {
            compile: {
                files: { 'build/all.min.js': ['build/all.js'] }
            }
        },
        replace: {
            compile: {
                options: {
                    patterns: [
                            {
                                match: /`/mg,
                                replacement: ''
                            },
                            {
                                match: /\/\*#\s*sourceMappingURL=.*?\s*\/\n/mg,
                                replacement: ''
                            }
                        ]
                    },
                    files: [ {expand: true, flatten: true, src: ['build/index.html'], dest: 'dist/'} ]
            }
        }
    });
  
    grunt.loadNpmTasks('grunt-contrib-pug');
    grunt.loadNpmTasks('grunt-contrib-sass');
    grunt.loadNpmTasks('grunt-contrib-concat');
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-html-build');
    grunt.loadNpmTasks('grunt-contrib-htmlmin');
    grunt.loadNpmTasks('grunt-replace');
  
    grunt.registerTask('compile', [
        'pug',
        'sass',
        'concat',
        'uglify',
        'htmlbuild',
        'htmlmin',
        'replace',
    ]);
};