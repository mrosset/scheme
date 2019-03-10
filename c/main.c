#include <libguile.h>

int main() {
        scm_init_guile();
        scm_shell(0, NULL);
}
