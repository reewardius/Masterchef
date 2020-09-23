module.exports = function(grunt) {
    // TODO: Delete bootstrap map
    // TODO: Minimize js files inside js folder

    grunt.initConfig({
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
        htmlbuild: {
            compile: {
                options: {
                    styles: {
                        main: 'build/style.min.css',
                        bootstrap: 'node_modules/bootstrap/dist/css/bootstrap.min.css',
                    },
                    scripts: {
                        main: 'src/js/main.js',
                        jquery: 'node_modules/jquery/dist/jquery.min.js',
                        sortable: 'node_modules/sortablejs/dist/sortable.umd.js',
                    }
                },
                src: 'build/index.html',
                dest: 'dist/'
            }
        },
        htmlmin: {
            compile: {
                options: {
                    removeComments: true,
                    collapseWhitespace: true
                },
                files: { 'dist/index.html': 'dist/index.html' }
            }
        },
    });
  
    grunt.loadNpmTasks('grunt-contrib-pug');
    grunt.loadNpmTasks('grunt-contrib-sass');
    grunt.loadNpmTasks('grunt-html-build');
    grunt.loadNpmTasks('grunt-contrib-htmlmin');

  
    grunt.registerTask('compile', ['pug', 'sass', 'htmlbuild', 'htmlmin']);
};