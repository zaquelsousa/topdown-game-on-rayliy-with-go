#version 330

uniform vec2 resolution;

void main() {
    vec3 cornflowerBlue = vec3(0.392, 0.584, 0.929);
    gl_FragColor = vec4(cornflowerBlue, 1.0);
}
