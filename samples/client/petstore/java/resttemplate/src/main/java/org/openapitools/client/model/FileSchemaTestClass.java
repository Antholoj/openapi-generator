/*
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.openapitools.client.model.ModelFile;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

/**
 * FileSchemaTestClass
 */
@JsonPropertyOrder({
  FileSchemaTestClass.JSON_PROPERTY_FILE,
  FileSchemaTestClass.JSON_PROPERTY_FILES
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", comments = "Generator version: 7.10.0-SNAPSHOT")

public class FileSchemaTestClass {
  public static final String JSON_PROPERTY_FILE = "file";
  @javax.annotation.Nullable  private ModelFile _file;

  public static final String JSON_PROPERTY_FILES = "files";
  @javax.annotation.Nullable  private List<ModelFile> files;

  public FileSchemaTestClass() {
  }

  /**
   * Constructor with all args parameters
   */
  public FileSchemaTestClass(@JsonProperty(JSON_PROPERTY_FILE) ModelFile _file, @JsonProperty(JSON_PROPERTY_FILES) List<ModelFile> files) {
    this._file = _file;
    this.files = files;
  }

  public FileSchemaTestClass _file(@javax.annotation.Nullable ModelFile _file) {
    
    this._file = _file;
    return this;
  }

  /**
   * Get _file
   * @return _file
   */
  @javax.annotation.Nullable  @JsonProperty(JSON_PROPERTY_FILE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public ModelFile getFile() {
    return _file;
  }


  @JsonProperty(JSON_PROPERTY_FILE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setFile(@javax.annotation.Nullable ModelFile _file) {
    this._file = _file;
  }

  public FileSchemaTestClass files(@javax.annotation.Nullable List<ModelFile> files) {
    
    this.files = files;
    return this;
  }

  public FileSchemaTestClass addFilesItem(ModelFile filesItem) {
    if (this.files == null) {
      this.files = new ArrayList<>();
    }
    this.files.add(filesItem);
    return this;
  }

  /**
   * Get files
   * @return files
   */
  @javax.annotation.Nullable  @JsonProperty(JSON_PROPERTY_FILES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<ModelFile> getFiles() {
    return files;
  }


  @JsonProperty(JSON_PROPERTY_FILES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setFiles(@javax.annotation.Nullable List<ModelFile> files) {
    this.files = files;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FileSchemaTestClass fileSchemaTestClass = (FileSchemaTestClass) o;
    return Objects.equals(this._file, fileSchemaTestClass._file) &&
        Objects.equals(this.files, fileSchemaTestClass.files);
  }

  @Override
  public int hashCode() {
    return Objects.hash(_file, files);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class FileSchemaTestClass {\n");
    sb.append("    _file: ").append(toIndentedString(_file)).append("\n");
    sb.append("    files: ").append(toIndentedString(files)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

  public static class Builder {
  
      private FileSchemaTestClass instance;
  
      public Builder() {
        this(new FileSchemaTestClass());
      }
  
      protected Builder(FileSchemaTestClass instance) {
        this.instance = instance;
      }
  
      public FileSchemaTestClass.Builder _file(ModelFile _file) {
          this.instance._file = _file;
        return this;
      }
        public FileSchemaTestClass.Builder files(List<ModelFile> files) {
          this.instance.files = files;
        return this;
      }
    
  
      /**
      * returns a built FileSchemaTestClass instance.
      *
      * The builder is not reusable.
      */
      public FileSchemaTestClass build() {
        try {
          return this.instance;
        } finally {
          // ensure that this.instance is not reused
          this.instance = null;
        }
      }
  
      @Override
      public String toString() {
        return getClass() + "=(" + instance + ")";
      }
    }
  
    /**
    * Create a builder with no initialized field.
    */
    public static FileSchemaTestClass.Builder builder() {
      return new FileSchemaTestClass.Builder();
    }
  
    /**
    * Create a builder with a shallow copy of this instance.
    */
    public FileSchemaTestClass.Builder toBuilder() {
      return new FileSchemaTestClass.Builder()
        ._file(getFile())
        .files(getFiles());
    }

}

