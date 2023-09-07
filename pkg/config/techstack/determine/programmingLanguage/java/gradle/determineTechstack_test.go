package gradle

import (
	"github.com/stretchr/testify/assert"
	"github.com/testcoders/detestcoder/pkg/constants/project"
	"log"
	"os"
	"testing"
)

func setup() {
	createTempTestfile()
}

func createTempTestfile() {
	content := []byte(gradleFile)
	err := os.WriteFile("build.gradle", content, 0644)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
}

func teardown() {
	err := os.Remove("build.gradle")
	if err != nil {
		log.Fatalf("Failed to remove %v: %v", "build.gradle", err)
	}
}

func TestDetermineTechstack(t *testing.T) {
	setup()
	defer teardown()

	ts := DetermineTechstack()

	assert.NotNil(t, ts)
	assert.Equal(t, ts.Language.Name, project.JAVA)
	assert.Equal(t, ts.Language.Version, "11")
	assert.Equal(t, ts.DependencyManager.Name, project.GRADLE)
	assert.Equal(t, ts.DependencyManager.Version, "")
	assert.Equal(t, ts.Framework.Name, "jakarta.ejb-api")
	assert.Equal(t, ts.Framework.Version, "4.0.0")
	assert.Equal(t, ts.TestDependencies[0].Name, "junit-jupiter-api")
	assert.Equal(t, ts.TestDependencies[0].Version, "5.9.2")
	assert.Equal(t, ts.TestDependencies[1].Name, "junit-jupiter-engine")
	assert.Equal(t, ts.TestDependencies[1].Version, "5.9.2")
	assert.Equal(t, ts.TestDependencies[2].Name, "junit-jupiter-engine")
	assert.Equal(t, ts.TestDependencies[2].Version, "5.9.2")
	assert.Equal(t, ts.TestDependencies[3].Name, "arquillian-junit5-container")
	assert.Equal(t, ts.TestDependencies[3].Version, "1.7.0.Final")
	assert.Equal(t, ts.TestDependencies[4].Name, "arquillian-junit5-core")
	assert.Equal(t, ts.TestDependencies[4].Version, "1.7.0.Final")
	assert.Equal(t, ts.TestDependencies[5].Name, "wildfly-arquillian-container-embedded")
	assert.Equal(t, ts.TestDependencies[5].Version, "5.0.1.Final")
	assert.Equal(t, ts.TestDependencies[6].Name, "mockito-junit-jupiter")
	assert.Equal(t, ts.TestDependencies[6].Version, "5.5.0")
}

var gradleFile = `
plugins {
    id 'java'
    id 'war'
}

group 'com.example'
version '1.0-SNAPSHOT'

repositories {
    mavenCentral()
}

ext {
    junitVersion = '5.9.2'
}

sourceCompatibility = '11'
targetCompatibility = '11'

tasks.withType(JavaCompile) {
    options.encoding = 'UTF-8'
}

dependencies {
    compileOnly('jakarta.ejb:jakarta.ejb-api:4.0.0')
    compileOnly('jakarta.servlet:jakarta.servlet-api:5.0.0')
    compileOnly('jakarta.transaction:jakarta.transaction-api:2.0.0')

    compileOnly 'org.projectlombok:lombok:1.18.28'
    annotationProcessor 'org.projectlombok:lombok:1.18.28'

    testImplementation("org.junit.jupiter:junit-jupiter-api:${junitVersion}")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:${junitVersion}")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:${junitVersion}")

    implementation platform('org.jboss.arquillian:arquillian-bom:1.7.0.Final')

    testImplementation 'org.jboss.arquillian.junit5:arquillian-junit5-container:1.7.0.Final'
    testImplementation 'org.jboss.arquillian.junit5:arquillian-junit5-core:1.7.0.Final'

    testImplementation 'org.wildfly.arquillian:wildfly-arquillian-container-embedded:5.0.1.Final'

    testImplementation group: 'org.mockito', name: 'mockito-junit-jupiter', version: '5.5.0'

}

test {
    useJUnitPlatform()
}
`
