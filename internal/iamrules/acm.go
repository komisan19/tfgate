package iamrules

func init() {
	// aws_acm_certificate
	register("aws_acm_certificate", OpCreate, Rule{
		BaseActions: []string{
			"acm:RequestCertificate",
			"acm:DescribeCertificate",
			"acm:AddTagsToCertificate",
		},
	})
	register("aws_acm_certificate", OpDelete, Rule{
		BaseActions: []string{
			"acm:DeleteCertificate",
			"acm:DescribeCertificate",
		},
	})
	register("aws_acm_certificate", OpUpdate, Rule{
		UpdateActions: []string{"acm:DescribeCertificate"},
		ConditionalActions: map[string][]string{
			"options": {"acm:UpdateCertificateOptions"},
			"tags":    {"acm:AddTagsToCertificate", "acm:RemoveTagsFromCertificate"},
		},
	})

	// aws_acm_certificate_validation
	register("aws_acm_certificate_validation", OpCreate, Rule{
		BaseActions: []string{
			"acm:DescribeCertificate",
			"route53:ChangeResourceRecordSets",
			"route53:GetChange",
		},
	})
	register("aws_acm_certificate_validation", OpDelete, Rule{
		BaseActions: []string{"acm:DescribeCertificate"},
	})
}
