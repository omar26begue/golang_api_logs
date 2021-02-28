package models

type Empresas struct {
	Identificador string `json:"identificador" gorm:"primary_key"`
	NombreEmpresa string `json:"nombre_empresa" gorm:"not null;comment:Nombre identificativo para la empresa"`
	Activo        bool   `json:"activo" gorm:"type:boolean"`
	IdOsde        string `json:"id_osde" gorm:"type:uuid"`
	Osde          Osde   `json:",omitempty" gorm:"ForeignKey:IdOsde;AssociationForeignKey:Identificador"`
}

type EmpresasRequest struct {
	NombreEmpresa string `json:"nombre_empresa" gorm:"not null;comment:Nombre identificativo para la empresa"`
	IdOsde        string `json:"id_osde" gorm:"type:uuid"`
	Activo        bool   `json:"activo" gorm:"type:boolean"`
}

type EmpresasResponse struct {
	Identificador string `json:"identificador" gorm:"primary_key"`
	NombreEmpresa string `json:"nombre_empresa" gorm:"not null;comment:Nombre identificativo para la empresa"`
	IdOsde        string `json:"id_osde" gorm:"type:uuid"`
	Activo        bool   `json:"activo" gorm:"type:boolean"`
}
