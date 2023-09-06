package maven

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
	content := []byte(pomString)
	err := os.WriteFile("pom.xml", content, 0644)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
}

func teardown() {
	err := os.Remove("pom.xml")
	if err != nil {
		log.Fatalf("Failed to remove %v: %v", "pom.xml", err)
	}
}

func TestDetermineTechstack(t *testing.T) {
	setup()
	defer teardown()

	ts := DetermineTechstack()

	assert.NotNil(t, ts)
	assert.Equal(t, ts.Language.Name, project.KOTLIN)
	assert.Equal(t, ts.Language.Version, "1.9.10")
	assert.Equal(t, ts.DependencyManager.Name, project.MAVEN)
	assert.Equal(t, ts.DependencyManager.Version, "")
	assert.Equal(t, ts.TestDependencies[0].Name, "assertj-core")
	assert.Equal(t, ts.TestDependencies[0].Version, "3.24.2")
	assert.Equal(t, ts.TestDependencies[1].Name, "spring-boot-starter-test")
	assert.Equal(t, ts.TestDependencies[1].Version, "")
	assert.Equal(t, ts.TestDependencies[2].Name, "kotlin-test")
	assert.Equal(t, ts.TestDependencies[2].Version, "1.9.10")
}

var pomString = `<?pom version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>org.m105.webshop</groupId>
        <artifactId>parent</artifactId>
        <version>0.0.1-SNAPSHOT</version>
        <relativePath>../pom.xml</relativePath>
    </parent>

    <artifactId>application</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <name>application</name>
    <description>application</description>

    <properties>
        <java.version>19</java.version>
        <kotlin.version>1.9.10</kotlin.version>
    </properties>

    <dependencies>
        <dependency>
            <groupId>org.bouncycastle</groupId>
            <artifactId>bcprov-jdk15on</artifactId>
            <version>1.68</version>
        </dependency>

        <dependency>
            <groupId>org.assertj</groupId>
            <artifactId>assertj-core</artifactId>
            <version>3.24.2</version>
            <scope>test</scope>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>org.jetbrains.kotlin</groupId>
            <artifactId>kotlin-stdlib-jdk8</artifactId>
            <version>${kotlin.version}</version>
        </dependency>
        <dependency>
            <groupId>org.jetbrains.kotlin</groupId>
            <artifactId>kotlin-test</artifactId>
            <version>${kotlin.version}</version>
            <scope>test</scope>
        </dependency>
    </dependencies>

    <build>
        <sourceDirectory>${project.basedir}/src/main/kotlin</sourceDirectory>
        <testSourceDirectory>src/test/kotlin</testSourceDirectory>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
            <plugin>
                <groupId>org.jetbrains.kotlin</groupId>
                <artifactId>kotlin-maven-plugin</artifactId>
                <version>${kotlin.version}</version>
                <executions>
                    <execution>
                        <id>compile</id>
                        <phase>compile</phase>
                        <goals>
                            <goal>compile</goal>
                        </goals>
                        <configuration>
                            <sourceDirs>
                                <source>src/main/kotlin</source>
                                <source>target/generated-sources/annotations</source>
                            </sourceDirs>
                        </configuration>
                    </execution>
                    <execution>
                        <id>test-compile</id>
                        <phase>test-compile</phase>
                        <goals>
                            <goal>test-compile</goal>
                        </goals>
                    </execution>
                </executions>
                <configuration>
                    <args>
                        <arg>-Xjsr305=strict</arg>
                    </args>
                    <compilerPlugins>
                        <plugin>spring</plugin>
                    </compilerPlugins>
                    <jvmTarget>1.8</jvmTarget>
                </configuration>
                <dependencies>
                    <dependency>
                        <groupId>org.jetbrains.kotlin</groupId>
                        <artifactId>kotlin-maven-allopen</artifactId>
                        <version>${kotlin.version}</version>
                    </dependency>
                </dependencies>
            </plugin>
        </plugins>
    </build>

</project>
`
