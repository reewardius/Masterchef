module.exports = function(grunt) {
    // TODO: Import FontAwesome css & font
    // TODO: Compile js files inside js folder

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
                        bootstrap: 'src/vendor/bootstrap/css/bootstrap.min.css'
                    },
                    scripts: {
                        main: 'src/js/main.js',
                        jquery: 'src/vendor/jquery/js/jquery.min.js',
                        sortable: 'src/vendor/sortable/js/sortable.min.js'
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